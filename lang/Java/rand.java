
import java.io.*;
import java.util.*;

import AdvReader;

public class rand {

  public static void test1(){
    Random rr;
    rr = new Random();

    System.out.println("\n== Test 1 ===");
    for(int ii = 0; ii < 10; ii++)
      System.out.println(rr.nextGaussian());
    }

  public static void test2(int maxGen, double mean, double scale){
    System.out.println("\n== Test 2 ===");
    for(int ii = 0; ii < maxGen; ii++)
      System.out.println(gaussianRandom(mean, scale));
    }

  public static double gaussianRandom(double mean, double scale){
    Random rr;
    rr = new Random();
    return rr.nextGaussian() * scale + mean;
    }

  public static void main(String[] options) {
    rand.test1();
    rand.test2(10, 5, 2);
    rand.test2(20, 5, 1);
  }
}

