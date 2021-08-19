package dcl;

public class Singleton {
    //volatile实例,保证线程可见性以及初始化的原子性
    private volatile static Singleton INSTANCE;
    //构造函数必须是私有的
    private Singleton (){
    }
    //双重校验锁单例调用
    public static Singleton getInstance() {
        //第一次判空提升效率
        if (INSTANCE == null) {
            synchronized (Singleton.class) {
                //第二次判空,避免重复初始化
                if (INSTANCE == null) {
                    INSTANCE = new Singleton();
                }
            }
        }
        return INSTANCE;
    }
}

