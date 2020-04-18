package com.alibaba.study.proxy;

public class JavaDeveloper implements  IDeveloper{
    private String name;

    public JavaDeveloper(String name) {
        this.name = name;
    }

    @Override
    public void code() {
        System.out.println(this.name + "is coding java");
    }

    @Override
    public void debug() {
        System.out.println(this.name + "is debuging java");
    }
}
