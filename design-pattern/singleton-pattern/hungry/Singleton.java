package hungry;

public class Singleton {
    //实例,使用类加载初始化,所以是线程安全
    private static final Singleton INSTANCE = new Singleton();
    //构造函数必须是私有的
    private Singleton (){
    }
    //单例调用
    public static Singleton getInstance() {
        return INSTANCE;
    }
}

