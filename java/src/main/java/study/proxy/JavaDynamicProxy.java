package study.proxy;


import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;


public class JavaDynamicProxy {

    public static void main(String[] args) {
        JavaDeveloper zack = new JavaDeveloper("zack");
        IDeveloper developer = (IDeveloper)Proxy.newProxyInstance(zack.getClass().getClassLoader(), zack.getClass().getInterfaces(), new InvocationHandler() {
                    @Override
                    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                        if (method.getName().equals("code")) {
                            System.out.println("Zack is preying before coding!");
                            return method.invoke(zack, args);
                        }
                        if (method.getName().equals("debug")) {
                            System.out.println("zack has no bug");
                            return null;
                        }

                        return null;
                    }
                }
        );

        developer.code();
        developer.debug();

    }


}
