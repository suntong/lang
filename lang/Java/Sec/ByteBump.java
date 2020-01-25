//**************************************
//     
// Name: Byte Bump Encode
//
// Description:My code takes a string, turns it to binary (a series of 1's and
// 0 's) then adds a few extra fake bits at the begining of the series and a
// couple at the end (to even it out). Then turns it back into a Readable
// string again. It also does the reverse process.
//
// By: John McDonald
//
// Inputs:A String
//
// Returns:An encoded String
//
// Assumes:This only encodes a string if
//     all letters are in the ascii character s
//     et. 0-255
//
// Side Effects:None known
//
//This code is copyrighted and has// limited warranties.Please see http://
//     www.Planet-Source-Code.com/xq/ASP/txtCod
//     eId.2732/lngWId.2/qx/vb/scripts/ShowCode
//     .htm//for details.
//     
//**************************************

public class ByteBump { 

  public static void main(String args[]){
    // sample use
    String coded = ByteBump.doByteBump("This is a sample phrase to see what it would look like");
    System.out.println(coded);
    System.out.println(ByteBump.undoByteBump(coded));
    }
                                  


  public static String doByteBump(String text){
    // change the text to binary, bump the bits down 6 notches, then turn it
    // back to text
    String bits = toBinary(text);
    String fakeBits = "";
    String rv = "";
                                                  
    // make a fake byte


    for(int i = 0; i < 8; i++){
      fakeBits += String.valueOf((int)(Math.rint(Math.random())));
      }
                                                  
    // put 6 fake bits from our fake byte on the front of the real bits
    // and the other 2 fake bits on the end
    bits = fakeBits.substring(0, 6) + bits + fakeBits.substring(6);
                                                  
    rv = fromBinary(bits);
                                                  
    return rv;
    }
                                          


  public static String undoByteBump(String text){
    // change the text to binary, bump the bits down 6 notches, then turn it
    // back to text
    if(text.length() == 1){return text;}
                                                          
                                                          
    String bits = toBinary(text);
    String rv = "";
                                                          
    // take off the fake bits that were added in the last process
    bits = bits.substring(6, bits.length() - 2);
                                                          
    rv = fromBinary(bits);
                                                          
    return rv;
    }
                                                  


  public static String toBinary(String text){
    // this turns a string into a sequence of bits
    String rv = "";
                                                          
    // go through every letter and turn it to binary
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
    return rv;
    }


  public static String fromBinary(String bits){
    // this turns a sequence of bits into a String
    String rv = "";
    // go through every 8 bits and turn then into a character


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
  }


