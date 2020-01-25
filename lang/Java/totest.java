import java.io.*;

// =========================================================== &c-beg ===
public class totest {

    private static void test1(String Option){
	File dirfile;

	dirfile = new File("/home/tong/discards/projects/bow-data");

	switch (Integer.parseInt(Option)) {
	case 1:
	    dirfile = new File("/home/tong/discards/projects/bow-data");
	    break;
	case 2:
	    dirfile = new File("/home/tong/discards/projects", "bow-data");
	    break;
	case 3:
	    dirfile = new File("/home/tong/discards/projects/bow-data", "");
	    break;
	default:
	    System.out.println("Invalid option");
	}
	System.out.println(dirfile);

	String files[];
	files = dirfile.list();
	System.out.println(files);

	for(int i=0; i<files.length; i++)
	    System.out.println(files[i]);
    }

    private static void test2(String thedir){
	File dirfile;

	dirfile = new File(thedir);
	System.out.println(dirfile);

	String files[];
	files = dirfile.list();
	System.out.println(files);

	for(int i=0; i<files.length; i++)
	    System.out.println(files[i]);
    }



// ---------------------------------------------------------- &c-main ---
    /**
     * Main method.
     */

    public static void main(String[] options) {
	//test1(options[1]);
	//System.out.println(Integer.parseInt("12"));
	test1("1");
	//	test1(new String("1"));
    }
}
