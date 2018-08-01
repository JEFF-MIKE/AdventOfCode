import java.util.*;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class DayOne{
    public static void main (String[] args) {
        List<String> lines = new ArrayList<String>();   
        BufferedReader br = null;
        try {
            br = new BufferedReader(new FileReader("d1.txt"));
            String line;
            while ((line = br.readLine()) != null) {
                //System.out.println(line);
                lines.add(line);
            }
        } catch (IOException e) {

        } finally {
            try {
            br.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
        System.out.println(lines);
        //private int blocks = 0;
        /*
        Directions are: left,up,down,right */
        //private String direction = "";
    }

}