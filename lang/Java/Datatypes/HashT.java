import java.util.Hashtable;
import java.util.HashMap;

public class HashT {
  public static final String K_ONE = "one";
  public static final String K_TWO = "two";
  public static final String K_THREE = "three";

  public static void TableTest() {

    Hashtable numbers = new Hashtable();
    numbers.put(K_ONE, new Integer(1));
    numbers.put(K_TWO, new Integer(2));
    numbers.put(K_THREE, new Integer(3));

    Integer n = (Integer)numbers.get(K_TWO);
    if (n != null) {
      System.out.println("two = " + n);
      }

    System.out.println("Hash contains key '" + K_ONE + "': " +
		       numbers.containsKey(K_ONE));
    System.out.println("Hash contains value '3': " +
		       numbers.containsValue(new Integer(3)));

    System.out.println("Hash content: " + numbers.toString());
    numbers.remove(K_TWO);
    System.out.println("Hash content: " + numbers.toString());
    
    };

  public static void MapTest() {

    HashMap numbers = new HashMap();
    numbers.put(K_ONE, new Integer(1));
    numbers.put(K_TWO, new Integer(2));
    numbers.put(K_THREE, new Integer(3));

    Integer n = (Integer)numbers.get(K_TWO);
    if (n != null) {
      System.out.println("two = " + n);
      }

    System.out.println("Hash contains key '" + K_ONE + "': " +
		       numbers.containsKey(K_ONE));
    System.out.println("Hash contains value '3': " +
		       numbers.containsValue(new Integer(3)));

    System.out.println("Hash content: " + numbers.toString());
    numbers.remove(K_TWO);
    System.out.println("Hash content: " + numbers.toString());

    };

  public static void main(String[] args) {
    System.out.println("*** Hash Table Test");
    TableTest();
    System.out.println("\n*** Hash Map Test");
    MapTest();
    }

  }
