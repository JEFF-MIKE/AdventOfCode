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
        // works
        List<String> splitLines = new ArrayList<String>();
        int listSize = lines.size();
        for(int i = 0;i<listSize;i++){
            String directions[] = lines.get(i).split("\\s");
            for (String element: directions){
                splitLines.add(element.replace(",",""));
            }
        }
        String direction = "up";
        int xvalue = 0;
        int yvalue = 0;
        for (String element: splitLines){
            int valueToAdd = Integer.parseInt(element.substring(1));
            switch(direction){
                case "up":
                if (element.charAt(0) == 'R'){
                    direction = "right";
                    xvalue += valueToAdd;
                    break;
                } else {
                    direction = "left";
                    xvalue -= valueToAdd;
                    break;
                }
                case "down":
                if (element.charAt(0) == 'R'){
                    direction = "left";
                    xvalue -= valueToAdd;
                    break;
                } else {
                    direction = "right";
                    xvalue += valueToAdd;
                    break;
                }
                case "left":
                if (element.charAt(0) == 'R'){
                    direction = "up";
                    yvalue += valueToAdd;
                    break;
                } else {
                    direction = "down";
                    yvalue -= valueToAdd;
                    break;
                }
                case "right":
                if (element.charAt(0) == 'R'){
                    direction = "down";
                    yvalue -= valueToAdd;
                    break;
                } else {
                    direction = "up";
                    yvalue += valueToAdd;
                    break;
                }
            }
        }
        System.out.println(Math.abs(xvalue) + Math.abs(yvalue));
    }
}