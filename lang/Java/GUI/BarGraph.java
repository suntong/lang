import java.awt.Container;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.GridLayout;
import java.awt.Rectangle;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;
import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JTextArea;
import javax.swing.JTextField;
import javax.swing.border.EtchedBorder;
import javax.swing.border.TitledBorder;


public class BarGraph {
  public static void main(String[] args) {
    BankFrame frame = new BankFrame();
    frame.setTitle("SavingAccount");
    frame.show();
    }
  }

class BankFrame extends JFrame {

  private JTextField amountInputArea;
  private JTextField rateInputArea;
  private JTextField yearInputArea;
	
	
  private JTextField resultField;
  private BarPanel barPanel;
  private JButton calcButton;
  private double balance;
  private double rate = 0.05;
  private double amount = 1000;
  private int year =4;
  private double[] data;
  private int MAXUNITS = 200;
  private Rectangle sample;

  public BankFrame() {
    setSize(300,600);
		
    addWindowListener(new WindowCloser());
		
    amountInputArea = new JTextField(20);
    rateInputArea = new JTextField(20);
    yearInputArea = new JTextField(20);
		
    resultField = new JTextField(20);
    resultField.setEditable(false);
		
    calcButton = new JButton("Calculate");
    calcButton.addActionListener(new ButtonListener());
		
    Container contentPane = getContentPane();
		
    barPanel = new BarPanel();
    barPanel.setBorder(new TitledBorder(new EtchedBorder(),"Bar"));
    contentPane.add(barPanel,"Center");
		
    JPanel amountPanel = new JPanel();
    amountPanel.add(new JLabel(" Amount:"));
    amountPanel.setBorder(new TitledBorder(new EtchedBorder(),"Amount"));
    amountPanel.add(amountInputArea);
		
    JPanel ratePanel = new JPanel();
    ratePanel.add(new JLabel("    Rate:"));
    ratePanel.setBorder(new TitledBorder(new EtchedBorder(),"Rate"));
    ratePanel.add(rateInputArea);
		
    JPanel yearPanel = new JPanel();
    yearPanel.add(new JLabel("    Year:"));
    yearPanel.setBorder(new TitledBorder(new EtchedBorder(),"Year"));
    yearPanel.add(yearInputArea);
		
		
    JPanel resultPanel = new JPanel();
    resultPanel.add(new JLabel("Balance:"));
    resultPanel.setBorder(new TitledBorder(new EtchedBorder(),"Balance"));
    resultPanel.add(resultField);
		
    JPanel buttonPanel = new JPanel();
    buttonPanel.add(calcButton);
		
    JPanel southPanel = new JPanel();
    southPanel.setLayout(new GridLayout(5,1));
    southPanel.add(amountPanel);
    southPanel.add(ratePanel);
    southPanel.add(yearPanel);
    southPanel.add(resultPanel);
    southPanel.add(buttonPanel);
		
    contentPane.add(southPanel,"South");
    }
	
  private class BarPanel extends JPanel {
	
    public void paintComponent(Graphics g) {
      super.paintComponent(g);
      Graphics2D g2 = (Graphics2D)g;
      double windowHeight = barPanel.getHeight();
      double unitHeight = barPanel.getHeight()/MAXUNITS;
      double unitWidth = getWidth()/year;
      int fixedYear = year + 1;
      double[] data = new double[fixedYear];
      for(int i = 1 ; i < fixedYear; i++) {
	data[i] = amount * Math.exp(rate * i);
	}
      for (int i = 1 ; i < fixedYear; i++) {
      	Rectangle sample = new
	  Rectangle((int)(i*unitWidth),(int)(barPanel.getHeight()-data[i]*unitHeight-1),(int)unitWidth,(int)(data[i]*unitHeight));
	sample = new Rectangle(10, 10, i* 10, i* 10);
      	g2.draw(sample);
	}
      }
    }
	
	
  private class ButtonListener implements ActionListener {
    
    /*	public void init() { 
	addActionListener(this);
	}
    */	
    public void actionPerformed(ActionEvent event) {
      amount = Double.parseDouble(amountInputArea.getText());
      rate = Double.parseDouble(rateInputArea.getText());
      year = Integer.parseInt(yearInputArea.getText());
			
      double balance = getBalance(amount, rate, year);
      resultField.setText("" + balance);
      repaint();
      }
    }
	
  public static double getBalance(double amount, double rate, double year) {
    double sum = amount * Math.exp(rate * year);
    return sum;
    }
	
  private class WindowCloser extends WindowAdapter {
    public void windowClosing(WindowEvent event) {
      System.exit(0);
      }
    }
	
  }






