/*
   Copyright (c) 2021 ffactory.org
   pretree is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
   EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
   MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package pretree

import (
	"maps"
	"testing"
)

func Test_Match(t *testing.T) {
	// 测试数据data包括 http请求方法，路由规则，客户端请求路径
	data := [][]string{
		{"POST", "/pet/{petId}/uploadImage", "/pet/12121/uploadImage"},
		{"POST", "/pet", "/pet"},
		{"PUT", "/pet", "/pet"},
		{"GET", "/pet/findByStatus", "/pet/findByStatus"},
		{"GET", "/pet/{petId}", "/pet/113"},
		{"GET", "/pet/{petId}/info", "/pet/12121/info"},
		{"POST", "/pet/{petId}", "/pet/12121"},
		{"DELETE", "/pet/{petId}", "/pet/12121"},
		{"GET", "/store/inventory", "/store/inventory"},
		{"POST", "/store/order", "/store/order"},
		{"GET", "/store/order/{orderId}", "/store/order/939"},
		{"DELETE", "/store/order/{orderId}", "/store/order/939"},
		{"POST", "/user/createWithList", "/user/createWithList"},
		{"GET", "/user/{username}", "/user/1002"},
		{"PUT", "/user/{username}", "/user/1002"},
		{"DELETE", "/user/{username}", "/user/1002"},
		{"GET", "/user/login", "/user/login"},
		{"GET", "/user/logout", "/user/logout"},
		{"POST", "/user/createWithArray", "/user/createWithArray"},
		{"POST", "/user", "/user"},
	}

	p := NewPreTree()
	for _, v := range data {
		method := v[0]
		sourceRule := v[1]
		p.Store(method, sourceRule)
	}

	for _, v := range data {
		method := v[0]
		urlPath := v[2]
		sourceRule := v[1]
		ok, rule, vars := p.Query(method, urlPath)
		if ok && rule == sourceRule {
			t.Logf("urlPath:%s match rule:%s result: %t vars: %s", urlPath, rule, ok, vars)
		} else {
			t.Errorf("method: %s urlPath:%s match rule:%s result: %t", method, urlPath, sourceRule, ok)
		}
	}
}

func Test_PathValues(t *testing.T) {
	// 测试数据data包括 http请求方法，路由规则，客户端请求路径
	data := [][]string{
		{"POST", "/pet/{petId}/uploadImage", "/pet/113/uploadImage"},
		{"GET", "/pet/{petId}", "/pet/113"},
		{"GET", "/pet/{petId}/info", "/pet/113/info"},
		{"POST", "/pet/{petId}", "/pet/113"},
		{"DELETE", "/pet/{petId}", "/pet/113"},
		{"GET", "/store/order/{orderId}", "/store/order/939"},
		{"DELETE", "/store/order/{orderId}", "/store/order/939"},
		{"GET", "/user/{username}", "/user/1002"},
		{"PUT", "/user/{username}", "/user/1002"},
		{"DELETE", "/user/{username}", "/user/1002"},
	}

	wants := []map[string]string{
		map[string]string{"petId": "113"},
		map[string]string{"petId": "113"},
		map[string]string{"petId": "113"},
		map[string]string{"petId": "113"},
		map[string]string{"petId": "113"},
		map[string]string{"orderId": "939"},
		map[string]string{"orderId": "939"},
		map[string]string{"username": "1002"},
		map[string]string{"username": "1002"},
		map[string]string{"username": "1002"},
	}

	p := NewPreTree()
	for _, v := range data {
		method := v[0]
		sourceRule := v[1]
		p.Store(method, sourceRule)
	}

	for i, v := range data {
		urlPath := v[2]
		want := wants[i]
		ok, got := p.PathValues(urlPath)
		if ok && maps.Equal(got, want) {
			t.Logf("urlPath:%s, got %v, wanted %v\n", urlPath, got, want)
		} else {
			t.Errorf("urlPath:%s, got %v, wanted %v\n", urlPath, got, want)
		}
	}
}

func Test_Build(t *testing.T) {
	// 测试数据data包括 http请求方法，路由规则，客户端请求路径
	data := [][]string{
		{"POST", "/pet/{petId}/uploadImage", "/pet/113/uploadImage"},
		{"POST", "/pet", "/pet"},
		{"PUT", "/pet", "/pet"},
		{"GET", "/pet/findByStatus", "/pet/findByStatus"},
		{"GET", "/pet/{petId}", "/pet/113"},
		{"GET", "/pet/{petId}/info", "/pet/113/info"},
		{"POST", "/pet/{petId}", "/pet/113"},
		{"DELETE", "/pet/{petId}", "/pet/113"},
		{"GET", "/store/inventory", "/store/inventory"},
		{"POST", "/store/order", "/store/order"},
		{"GET", "/store/order/{orderId}", "/store/order/939"},
		{"DELETE", "/store/order/{orderId}", "/store/order/939"},
		{"POST", "/user/createWithList", "/user/createWithList"},
		{"GET", "/user/{username}", "/user/1002"},
		{"PUT", "/user/{username}", "/user/1002"},
		{"DELETE", "/user/{username}", "/user/1002"},
		{"GET", "/user/login", "/user/login"},
		{"GET", "/user/logout", "/user/logout"},
		{"POST", "/user/createWithArray", "/user/createWithArray"},
		{"POST", "/user", "/user"},
	}

	vars := map[string]string{
		"username": "1002",
		"petId":    "113",
		"orderId":  "939",
	}

	for _, v := range data {
		sourceRule := v[1]
		want := v[2]
		got := SetPathValues(sourceRule, vars)
		if got != want {
			t.Errorf("got %v path, wanted %v\n", got, want)
		}
	}
}
