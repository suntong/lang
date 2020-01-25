import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.geom.Ellipse2D;
import javax.swing.JFrame;
import javax.swing.JPanel;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;


import java.awt.Container;

public class ClickEgg 
{
  public static void main(String[] args)
  {
    PictureFrame frame = new PictureFrame();
    frame.setTitle("ClickEgg");
    frame.setSize(300, 300);
    frame.show(); 
    }
  }

class PictureFrame extends JFrame
{
  private PicturePanel panel;
  private Ellipse2D.Double egg;

  public PictureFrame() {
    panel = new PicturePanel();
    Container contentPane = getContentPane();
    contentPane.add(panel,"Center");
		
    addWindowListener(new WindowCloser());

    egg = new Ellipse2D.Double(0,0, 30, 50);
    }
	
  private class PicturePanel extends JPanel {
		
    private class MouseClickListener extends MouseAdapter {
      public void mouseClicked(MouseEvent event) {
	int x = event.getX();
	int y = event.getY();
	egg.setFrame(x - 15,y - 15,30, 30);
	repaint();
	}
      }

    public PicturePanel(){
      MouseClickListener Mlistener = new MouseClickListener();
      addMouseListener(Mlistener);
      }

    public void paintComponent(Graphics g) {
      super.paintComponent(g);
      Graphics2D g2 = (Graphics2D)g;
      String message = "Please click";
      g2.drawString(message,50,100);
      g2.draw(egg);
      }
    }

  private class WindowCloser extends WindowAdapter
  {
    public void windowClosing(WindowEvent event)
    {
      System.exit(0);
      }
    }
  }

