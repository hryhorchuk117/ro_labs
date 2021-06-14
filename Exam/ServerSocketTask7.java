package Server;

import PhoneDTO;

import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;
import java.rmi.RemoteException;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

public class ServerSocketTask7 {
    private ServerSocket server = null;
    private Socket socket = null;
    private PrintWriter output = null;
    private BufferedReader in = null;
    private static final String split = "#";

    public void start(int port) throws IOException {
        server = new ServerSocket(port);
        while (true) {
            socket = server.accept();
            in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
            String query = in.readLine();
            output=new PrintWriter(socket.getOutputStream(), true);
            try (ObjectOutputStream out = new ObjectOutputStream( socket.getOutputStream() )) {
                new Thread( ( ) -> {
                    processQuery( query, out );
                } ).start( );
            }
        }
    }
    public ArrayList<PhoneDTO> findPhonesWhereTownCallsMoreThan(Integer value) throws RemoteException {
        ArrayList<PhoneDTO> res=new ArrayList<>(  );
        for ( PhoneDTO phone:
                Data.phones) {
            if(phone.getTownCalls()>value){
                res.add(phone);
            }
        }
        return res;
    }

    public ArrayList<PhoneDTO> findPhonesWhereOutTownCallsPresent() throws RemoteException {
        ArrayList<PhoneDTO> res=new ArrayList<>(  );
        for ( PhoneDTO phone:
                Data.phones) {
            if(phone.getOutOfTownCalls()>0){
                res.add(phone);
            }
        }
        return res;
    }

    public ArrayList<PhoneDTO> sortedAlphabetic() throws RemoteException {
        ArrayList<PhoneDTO> res=Data.phones;
        res.sort( Comparator.comparing( PhoneDTO::getFirstName ) );
        res.sort( Comparator.comparing( PhoneDTO::getLastname ) );
        res.sort( Comparator.comparing( PhoneDTO::getSecondName ) );
        return res;
    }
    private boolean processQuery(String query, ObjectOutputStream out) {
        try {
            if (query == null) {
                return false;
            }

            String[] fields = query.split(split);
            if (fields.length == 0) {
                return true;
            } else {
                var action = fields[0];

                switch (action) {
                    case "findPhonesWhereTownCallsMoreThan":
                        var value = Integer.parseInt(fields[1]);
                        out.writeObject(findPhonesWhereTownCallsMoreThan(value));
                        break;
                    case "findPhonesWhereOutTownCallsPresent":
                        out.writeObject(findPhonesWhereOutTownCallsPresent());
                        break;
                    case "sortedAlphabetic":
                        out.writeObject(sortedAlphabetic());
                        break;
                }
            }

            return true;
        } catch (IOException e) {
            return false;
        }
    }

    public static void main(String[] args) {
        try {
            ServerSocketTask7 server = new ServerSocketTask7();
            server.start(3000);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
