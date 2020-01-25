import java.io.*;
import java.math.BigInteger;
import java.security.*;
import java.security.spec.*;

public class SnGenerator {

  public static String encrypt(String text){
    // Do something very odd here to make an encrypted string
    String rv = "";
                                                          
    for(int i = 0; i < text.length(); i++){
      int curByte = (int)text.charAt(i);
      if(curByte >= 128){
	rv += 1;
	curByte -= 128;
	}
      else{rv += 0;}

      if(curByte >= 64){
	rv += 1;
	curByte -= 64;
	}
      else{rv += 0;}

      if(curByte >= 32){
	rv += 1;
	curByte -= 32;
	}
      else{rv += 0;}

      if(curByte >= 16){
	rv += 1;
	curByte -= 16;
	}
      else{rv += 0;}

      if(curByte >= 8){
	rv += 1;
	curByte -= 8;
	}
      else{rv += 0;}

      if(curByte >= 4){
	rv += 1;
	curByte -= 4;
	}
      else{rv += 0;}

      if(curByte >= 2){
	rv += 1;
	curByte -= 2;
	}
      else{rv += 0;}

      if(curByte >= 1){
	rv += 1;
	curByte -= 1;
	}
      else{rv += 0;}
      }

    String bits = rv;
    String fakeBits = "";
    for(int i = 0; i < 8; i++){
      fakeBits += String.valueOf((int)(Math.rint(Math.random())));
      }
    bits = fakeBits.substring(0, 6) + bits + fakeBits.substring(6);

    rv = new String("");
    for(int i = 0; i < bits.length(); i++){
      int curByte = 0;
      if(bits.charAt(i) == '1'){curByte += 128;}
      i++;
      if(bits.charAt(i) == '1'){curByte += 64;}
      i++;
      if(bits.charAt(i) == '1'){curByte += 32;}
      i++;
      if(bits.charAt(i) == '1'){curByte += 16;}
      i++;
      if(bits.charAt(i) == '1'){curByte += 8;}
      i++;
      if(bits.charAt(i) == '1'){curByte += 4;}
      i++;
      if(bits.charAt(i) == '1'){curByte += 2;}
      i++;
      if(bits.charAt(i) == '1'){curByte += 1;}
      rv += (char)curByte;
      }
    return rv;
    }

  public static void baOutputA(String baName, byte[] ba) {
    System.out.print("  private final byte[] " + baName + " = { (byte)0,");
    for(int i = 0; i < ba.length; ){
      if (i % 5 == 0) System.out.print("\n    ");
      System.out.print("(byte)0x" + Integer.toString(ba[i] & 0xff, 16));
      if (++i < ba.length) System.out.print(", ");
      }
    System.out.println("}; // " + baName);
    }

  public static String baOutput1(byte[] ba) {
    String rv = "";
    for(int i = 0; i < ba.length; i++){
      rv += Integer.toString(ba[i] & 0xff, 16).toUpperCase();
      }
    return rv;
    }

  public static void main(String[] args) {

    if (args.length != 1) {
      System.err.println("Usage: java SnGenerator <company name>");
      return;
    }

    String compName = args[0];

    if (compName.length() < 8 ) {
      System.err.println("Company name must be at least eight characters long");
    }
			 
    try {
      byte[] key = encrypt(compName).getBytes("ISO-8859-1");

      // use SHA digest algorithm
      MessageDigest sha = MessageDigest.getInstance("SHA");
      byte[] serialNumber = sha.digest(key);

      baOutputA("key", key);
      baOutputA("serialNumber", serialNumber);
      String ba1 = baOutput1(serialNumber);
      System.out.println("serial number =\n" + ba1);
      System.out.println(Long.toHexString(-1).toUpperCase());
      final byte[] serialNumber1 = { (byte)0,
    (byte)0xb9, (byte)0x6b, (byte)0xdf, (byte)0x8f, (byte)0x88, 
    (byte)0xd4, (byte)0xec, (byte)0x96, (byte)0x3b, (byte)0x69, 
    (byte)0x69, (byte)0xf, (byte)0xc6, (byte)0x82, (byte)0x6a, 
    (byte)0x3e, (byte)0x87, (byte)0x7f, (byte)0x6d, (byte)0x77}; // serialNumber

      BigInteger bi0 = new BigInteger(serialNumber);
      BigInteger bi1 = new BigInteger(serialNumber1);
      BigInteger bi12 = new BigInteger("B96BDF8F88D4EC963B69690FC6826A3E877F6D77", 16);
      BigInteger bi2 = new BigInteger(ba1, 16);
      System.out.println(bi0.equals(bi2) + bi0.toString(16)  + ", " + bi2.toString(16));
      System.out.println(bi1.equals(bi12) + bi1.toString(16)  + ", " + bi12.toString(16));


      }
    catch (Exception e) {
      System.err.println(e);
      e.printStackTrace();
    }
  }

}
