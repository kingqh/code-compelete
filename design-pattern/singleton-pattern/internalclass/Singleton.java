package internalclass;

public class Singleton {
    //静态内部类,只有显式调用才会加载
    private static class SingletonFactory {
        private static final Singleton INSTANCE = new Singleton();
    }
    //构造函数必须是私有的
    private Singleton (){
    }
    //单例调用
    public static Singleton getInstance() {
        return SingletonFactory.INSTANCE;
    }
}

