/*
 * Copyright 2011 Alibaba.com All right reserved. This software is the
 * confidential and proprietary information of Alibaba.com ("Confidential
 * Information"). You shall not disclose such Confidential Information and shall
 * use it only in accordance with the terms of the license agreement you entered
 * into with Alibaba.com.
 */
package study.rpc.test;
/**
 * HelloServiceImpl
 *
 * @author william.liangf
 */

public class HelloServiceImpl implements HelloService {
    public String hello(String name) {
        return "Hello " + name;
    }

    @Override
    public String echo(String str) {
        return str;
    }


}