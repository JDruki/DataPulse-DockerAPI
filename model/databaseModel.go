package model

import (
	"Auto/dao"
	"Auto/util"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 辅助函数-将interface转为对应字典类型
func mapToStringArray(m map[string]interface{}) []string {
	pairs := make([]string, 0, len(m))
	for k, v := range m {
		switch v := v.(type) {
		case string:
			fmt.Println("转换str")
			pairs = append(pairs, fmt.Sprintf("`%s` = '%s'", k, v))
		case int:
			fmt.Println("转换int")
			pairs = append(pairs, fmt.Sprintf("`%s` = %d", k, v))
		case float32, float64:
			fmt.Println("转换float")
			f := reflect.ValueOf(v).Float()
			pairs = append(pairs, fmt.Sprintf("`%s` = %.2f", k, f))
		case bool:
			fmt.Println("转换bool")
			b := strconv.FormatBool(v)
			pairs = append(pairs, fmt.Sprintf("`%s` = %s", k, b))
		case *string:
			fmt.Println("转换*str")
			s := *v
			pairs = append(pairs, fmt.Sprintf("`%s` = '%s'", k, s))
		case *int:
			fmt.Println("转换*int")
			i := *v
			pairs = append(pairs, fmt.Sprintf("`%s` = %d", k, i))
		case *float32, *float64:
			fmt.Println("转换*float")
			fv := reflect.ValueOf(v).Elem().Float()
			pairs = append(pairs, fmt.Sprintf("`%s` = %.2f", k, fv))
		case *bool:
			fmt.Println("转换*bool")
			bv := *v
			b := strconv.FormatBool(bv)
			pairs = append(pairs, fmt.Sprintf("`%s` = %s", k, b))
		case time.Time:
			fmt.Println("转换time.Time")
			t := v.Format(time.RFC3339)
			pairs = append(pairs, fmt.Sprintf("`%s` = '%s'", k, t))
		default:
			panic(fmt.Sprintf("Unsupported type for value: %T", v))
		}
	}
	return pairs
}

// SearchUserDataBase  通过用户查询日志
func SearchUserDataBase(value map[string]interface{}, table string) ([]map[string]interface{}, error) {
	var (
		query string
	)
	query = fmt.Sprintf("SELECT * FROM `%s` WHERE state != 0", table)

	wherePairs := mapToStringArray(value)
	fmt.Println(wherePairs)

	if len(wherePairs) > 0 {
		query += " AND " + strings.Join(wherePairs, " AND ")
	}
	fmt.Println("拼接query为" + query + "\n")
	err := dao.Db.Ping()
	if err != nil {
		fmt.Printf("数据库不健康 : %s\n", err.Error())
		return nil, err
	}
	// 执行查询
	rows, err := dao.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	logs := make([]map[string]interface{}, 0)

	// 迭代查询结果
	for rows.Next() {
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		// 扫描行数据
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		entry := make(map[string]interface{})
		// 将行数据存储在映射中
		for i, col := range values {
			if col != nil {
				entry[columns[i]] = col
			} else {
				entry[columns[i]] = nil
			}
		}
		logs = append(logs, entry)
	}
	//解码切片为utf-8格式
	if err = rows.Err(); err != nil {
		return nil, err
	}
	for _, logEntry := range logs {
		for key, value := range logEntry {
			// 检查值是否是字节切片类型
			if byteSlice, ok := value.([]byte); ok {
				// 将字节切片转换为字符串
				strValue := string(byteSlice)
				logEntry[key] = strValue
			}
		}
	}
	//fmt.Println(logs)
	return logs, nil
}

func SetDataBase(value map[string]interface{}, table string, ID int) error {
	var (
		query    string
		setPairs []string
	)
	logDate := util.GetUnixTime()
	setPairs = mapToStringArray(value)
	query = fmt.Sprintf("UPDATE `%s` SET `update_time` = %d,  %s WHERE id = %d", table, logDate, strings.Join(setPairs, ", "), ID)
	fmt.Println("拼接query为" + query + "\n")
	err := dao.Db.Ping()
	if err != nil {
		fmt.Printf("数据库不健康 : %s\n", err.Error())
		return err
	}
	// 执行更新操作
	_, err = dao.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InSetDataBase 插入新值
func InSetDataBase(value map[string]interface{}, table string) error {
	var (
		query string
	)
	logDate := util.GetUnixTime()
	state := 1
	query = fmt.Sprintf("INSERT INTO `%s` SET `state` = %d, `create_time` = %d", table, state, logDate)
	wherePairs := mapToStringArray(value)
	if len(wherePairs) > 0 {
		query += " ," + strings.Join(wherePairs, " ,")
	}
	query += ";"

	fmt.Println("拼接query为" + query + "\n")

	err := dao.Db.Ping()
	if err != nil {
		fmt.Printf("数据库不健康 : %s\n", err.Error())
		return err
	}
	_, err = dao.Db.Exec(query)

	return err
}

func DeleteDataBase(value map[string]interface{}, table string) error {
	var (
		query string
	)
	logDate := util.GetUnixTime()
	query = fmt.Sprintf("UPDATE %s SET `update_time` = %d, state = 0 WHERE state = 1", table, logDate)
	wherePairs := mapToStringArray(value)
	if len(wherePairs) > 0 {
		query += " AND " + strings.Join(wherePairs, " AND ")
	}
	fmt.Println("拼接query为" + query + "\n")
	err := dao.Db.Ping()
	if err != nil {
		fmt.Printf("数据库不健康 : %s\n", err.Error())
		return err
	}
	_, err = dao.Db.Exec(query)
	return err
}
