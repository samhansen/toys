import java.io.BufferedReader;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.lang.Integer;
import java.lang.String;
import java.lang.System;
import java.util.logging.Level;
import java.util.logging.Logger;

public class Spiral {
  private static final Logger logger = Logger.getLogger("Spiral.main");

  private static void spiral(int x1, int x2, int y1, int y2,
      int m, int n, String[][] s) {

    boolean top = x2 > x1 && y2 > y1;
    boolean right = top && x2 > x1 + 1;
    boolean bottom = right && y2 - 1 > y1;
    boolean left = bottom && x2 - 1 > x1;

    // Stop recursing only when there is nothing left to print.
    if (!top && !right && !bottom && !left) {
      return;
    }

    if (top) {
      for (int i=y1; i<y2; i++) {
        System.out.printf("%s ", s[x1][i]);
      }
    }

    if (right) {
      for (int i=x1+1; i<x2; i++) {
        System.out.printf("%s ", s[i][y2-1]);
      }
    }

    if (bottom) {
      for (int i=y2-2; i>=y1; i--) {
        System.out.printf("%s ", s[x2-1][i]);
      }
    }

    if (left) {
      for (int i=x2-2; i>=x1+1; i--) {
        System.out.printf("%s ", s[i][y1]);
      }
    }

    // Recurse into the subproblem created by shrinking each edge by 1 col/row.
    spiral(x1+1, x2-1, y1+1, y2-1, m, n, s);
  }

  public static void main(String[] args) {
    File fd = new File(args[0]);
    try(BufferedReader br = new BufferedReader(new FileReader(fd))) {
      for(String line; (line = br.readLine()) != null; ) {
        String[] parts = line.split("\\s*;\\s*");
        String[] values = parts[2].split("\\s+");
        int m = Integer.parseInt(parts[0]);
        int n = Integer.parseInt(parts[1]);
        String[][] s = new String[m][n];

        // Populate array.
        for (int i=0; i<m; i++) {
          for (int j=0; j<n; j++) {
            s[i][j] = values[(i*n) + j];
          }
        }

        spiral(0, m, 0, n, m, n, s);
        System.out.println();
      }
    } catch (FileNotFoundException ex) {
      logger.log(Level.SEVERE, ex.toString());
      System.exit(1);
    } catch (IOException ex) {
      logger.log(Level.SEVERE, ex.toString());
      System.exit(1);
    }
  }
}
