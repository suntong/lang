import javax.swing.JFrame;
import javax.swing.JPanel;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.Container;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.geom.Line2D;
import java.awt.geom.Point2D;
import java.awt.Rectangle;
import javax.swing.JTextField;
import java.awt.Color;


class ProgramFrame extends JFrame {
	
  private JTextField textField;
  private PicturePanel PPanel;
  private int SquareCount =1;
  int UNIT;
  int ColorLocationX = 0;
  int ColorLocationY = 0;
	
  public ProgramFrame() {
    setSize(300, 300);
		
    addWindowListener(new WindowCloser());
		
    PPanel = new PicturePanel();
		
    textField = new JTextField();
    textField.addActionListener(new TextFieldListener());
		
    Container contentPane = getContentPane();
    contentPane.add(textField,"North");
    contentPane.add(PPanel,"Center");
    }
	
  private class TextFieldListener implements ActionListener {
    public void actionPerformed(ActionEvent event) {
      String input = textField.getText();
      PPanel.setSquareCount(Integer.parseInt(input));
      textField.setText("");	
      }
    }
	
	
  private class PicturePanel extends JPanel {

    public PicturePanel(){
      MouseClickedListener Mlistener = new MouseClickedListener();
      addMouseListener(Mlistener);
      }

    public void paintComponent(Graphics g) {
      super.paintComponent(g);
      Graphics2D g2 = (Graphics2D)g;
      int n = SquareCount;
      int ScreenX = getWidth();
      int ScreenY = getHeight();
      UNIT = (ScreenY - 20)/n;
			
			
      for(int i = 0; i<n * UNIT; i+=UNIT) {
	  for( int j = 0;  j <n * UNIT; j+=UNIT) {
	      Rectangle Single = new Rectangle(i,j,UNIT,UNIT);
	      g2.draw(Single);
	      }
	  }

      if (ColorLocationX != 0) {
	  Rectangle Colored = new 
	    Rectangle(ColorLocationX,ColorLocationY,UNIT, UNIT);
	  g2.setColor(Color.black);
	  g2.fill(Colored);
	  }
			
      String message = "Please click on any single square you like.";
      g2.drawString(message,10,ScreenY - 10);
      }
		
    private class MouseClickedListener extends MouseAdapter {
      public void mouseClicked(MouseEvent event) {  
	int mouseLocationX = event.getX();
	int mouseLocationY = event.getY();
	ColorLocationX = (int)mouseLocationX/UNIT*UNIT;
	ColorLocationY = (int)mouseLocationY/UNIT*UNIT;
	repaint();
	}
      }
	
    public void setSquareCount(int count) {
      SquareCount = count;
      repaint();
      }
    }
	
  private class WindowCloser extends WindowAdapter {
    public void windowClosing(WindowEvent event) {
      System.exit(0);
      }
    }
  }

public class Grid {
  public static void main(String[] args) {
    ProgramFrame frame = new ProgramFrame();
    frame.setTitle("ColorSqure");
    frame.show();
    }
  }

