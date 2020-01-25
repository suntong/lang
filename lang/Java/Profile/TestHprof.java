package com.develop.demos;

import java.io.IOException;

public class TestHprof {
    public static String cat = null;
    public final static int loop=5000;
    
    public static void makeString() {
        cat = new String();
        for (int n=0; n<loop; n++) {
            addToCat("more");
        }    
    }
    
    public static void addToCat(String more) {
        cat = cat + more;
    }
    
    public static void makeStringInline() {
        cat = new String();
        for (int n=0; n<loop; n++) {
            cat = cat + "more";
        }
    }

    public static void makeStringWithLocal() {
        String tmp = new String();
        for (int n=0; n<loop; n++) {
           tmp = tmp + "more";
        }
        cat = tmp;
    }
    
    public static void makeStringWithBuffer() {
        StringBuffer sb = new StringBuffer();
        for (int n=0; n<loop; n++) {
            sb.append("more");
        }
        cat = sb.toString();
    }
    public static void main(String[] args) {
    	long begin = System.currentTimeMillis();
    	
    	if (null != System.getProperty("WaitForProfiler")) {
	    	System.out.println("Start your profiler, then press any key to begin...");
    		try {
    			System.in.read();
    		}
    		catch (IOException ioe) {
	    	}
    	}
            
        makeString();
        makeStringInline();
        makeStringWithLocal();
        makeStringWithBuffer();
        
        long end = System.currentTimeMillis();
        System.out.println("Total run time of " + (end - begin) + " milliseconds");
    }
} 
 
