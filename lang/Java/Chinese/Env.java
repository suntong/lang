/*
 * Copyright (c) 2002 Email: chedongATbigfoot.com/chedongATchedong.com
 * $Id: hello_unicode.html,v 1.6 2003/11/09 07:57:11 chedong Exp $
 */

import java.util.*;
import java.text.*;

/**
 * Goal:
 *     Show env vars and JVM default values
 * In:none
 * Out:
 *     1 supported LOCALEs
 *     2 JVM default values
 */

public class Env {
    /**
     *  main entrance
     */
    public static void main(String[] args) {
    	
        System.out.println("Hello, it's: " +  new Date());

        //print available locales
        Locale list[] = DateFormat.getAvailableLocales();
        System.out.println("======System available locales:======== ");
        for (int i = 0; i < list.length; i++) {
            System.out.println(list[i].toString() + "\t" + list[i].getDisplayName());
        }

        //print JVM default properties
        System.out.println("======System property======== ");
        System.getProperties().list(System.out);
    }
}
