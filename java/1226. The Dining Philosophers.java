class DiningPhilosophers {
    private Semaphore mutex;
    private Semaphore sema[];

    public DiningPhilosophers() {
        mutex = new Semaphore(1);
        sema = new Semaphore[]{
            new Semaphore(1),
            new Semaphore(1),
            new Semaphore(1),
            new Semaphore(1),
            new Semaphore(1)
        };
    }

    // call the run() method of any runnable to execute its code
    public void wantsToEat(int philosopher,
                           Runnable pickLeftFork,
                           Runnable pickRightFork,
                           Runnable eat,
                           Runnable putLeftFork,
                           Runnable putRightFork) throws InterruptedException {
        mutex.acquire(); // 同时只有一个philosopher在进食
        sema[philosopher].acquire();
        pickLeftFork.run();
        sema[(philosopher+1)%5].acquire();
        pickRightFork.run();

        eat.run();

        putLeftFork.run();
        sema[philosopher].release();
        putRightFork.run();
        sema[(philosopher+1)%5].release();
        mutex.release();
    }
}