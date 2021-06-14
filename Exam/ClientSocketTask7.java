package Exam;

import PhoneDTO;

import java.io.*;
import java.net.Socket;
import java.util.ArrayList;
import java.util.List;

public class ClientSocketTask7 {
    private final Socket socket;
    private final PrintWriter out;
    private final ObjectInputStream in;
    private static final String split = "#";

    ClientSocketTask7(String ip, int port) throws IOException {
        socket = new Socket(ip, port);
        in = new ObjectInputStream(socket.getInputStream( ));
        out = new PrintWriter(socket.getOutputStream( ), true);
    }
    
    public void disconnect() {
        try {
            socket.close( );
        } catch (IOException e) {
            e.printStackTrace( );
        }
    }

    public Object sendQuery(String query) throws IOException, ClassNotFoundException {
        String[] fields = query.split( split );
        out.println(query);
        ArrayList<PhoneDTO> phones = new ArrayList<>();
        var action = fields[0];
        phones = (ArrayList<PhoneDTO>) in.readObject();
        return phones;
    }

    public static void main(String[] args) throws IOException, ClassNotFoundException {
        ClientSocketTask7 client = new ClientSocketTask7( "localhost", 3001 );
        BufferedReader reader = new BufferedReader(
                new InputStreamReader( System.in ) );
        while (true) {
            String query = reader.readLine( );
            client.sendQuery( query );
        }
    }
}
