package study.blockingqueue;

import java.util.concurrent.BlockingDeque;
import java.util.concurrent.LinkedBlockingDeque;

public class MyBlockingQueue<T> {
    private BlockingDeque<T> blockingDeque = new LinkedBlockingDeque<>(50);

    public void add(T msg) {
        blockingDeque.add(msg);
    }

    public T remove() {
        T t = null;
        try {
             t = blockingDeque.take();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        return t;
    }

    public static void main(String[] args) {
        MyBlockingQueue<String> myBlockingQueue = new MyBlockingQueue<>();

        Thread producerThread = new Thread(new Runnable() {
            private int idx = 1;

            @Override
            public void run() {
                produce();
            }

            private void produce() {
                while (true) {
                    myBlockingQueue.add(String.valueOf(idx++));
                    System.out.println("add msg to queue: " + String.valueOf(idx));
                }
            }
        });

        Thread consumerThread = new Thread(new Runnable() {
            @Override
            public void run() {
                consume();
            }

            private void consume() {
                while (true) {
                    String msg = myBlockingQueue.remove();
                    System.out.println("get msg from queue: " + msg);
                }
            }
        });


        producerThread.start();
        consumerThread.start();


    }

}


