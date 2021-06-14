package Exam;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
@Data
@NoArgsConstructor
@AllArgsConstructor
public class PhoneDTO implements Serializable {
    private long id;
    private String firstName;
    private String lastname;
    private String secondName;
    private String address;
    private Integer creditCard;
    private Double debit;
    private Double credit;
    private Integer townCalls;
    private Integer outOfTownCalls;
}
