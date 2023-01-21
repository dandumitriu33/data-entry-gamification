import { Component } from '@angular/core';

@Component({
  selector: 'app-paper-sim',
  templateUrl: './paper-sim.component.html',
  styleUrls: ['./paper-sim.component.css']
})

export class PaperSimComponent {
  firstName = "Jim"
  lastName = "Smith"
  year = 1999
  make = "Honda"
  vin = "123ABC345"
  state = "NY"

  ngOnInit(): void {
    this.newReceiptData();
  }

  newReceiptData() {
    this.firstName = this.getRandomFirstName();
    this.lastName = this.getRandomLastName();
    this.year = this.getRandomYear();
    this.make = this.getRandomMake();
    this.vin = this.getRandomVIN();
    this.state = this.getRandomState();
  }

  getRandomFirstName() {
    const firstNames = ["James", "Robert", "Michael", "John", "David", "William", "Richard", 
                        "Joseph", "Thomas", "Charles", "Christopher", "Daniel", "Matthew",
                        "Anthony", "Mark", "Donald", "Steven", "Paul", "Andrew", "Kevin",
                        "Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara",
                        "Susan", "Jessica", "Sarah", "Karen", "Lisa", "Nancy", "Betty",
                        "Margaret", "Sandra", "Ashley", "Kimberly", "Emily", "Donna"];
    const random = Math.floor(Math.random() * firstNames.length);
    return firstNames[random]
  }

  getRandomLastName() {
    const lastNames = ["Smith", "Jones", "Williams", "Taylor", "Brown", "Davies", "Evans",
                        "Wilson", "Johnson", "Roberts", "Robinson", "Walker", "Hall",
                        "Lewis", "Harris", "Clarke", "Jackson", "Wood", "Turner",
                        "Martin", "Cooper", "Hill", "Morris", "Ward", "Moore", "Lee",
                        "Baker", "Allen", "Phillips", "Ford", "Day", "West"];
    const random = Math.floor(Math.random() * lastNames.length);
    return lastNames[random]
  }

  getRandomYear() {
    const years = [1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004,
                    2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014,
                    2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023];
    const random = Math.floor(Math.random() * years.length);
    return years[random]
  }

  getRandomMake() {
    const makes = ["Honda", "Toyota", "Chevrolet", "Ford", "Mercedes-Benz", "Jeep",
                    "BMW", "Porsche", "Subaru", "Nissan", "Cadillac", "VolksWagen",
                    "Lexus", "Audi", "Volvo", "Jaguar", "GMC", "Buick", "Acura",
                    "Dodge", "Hyundai", "Lincoln", "Mazda", "Tesla", "Kia",
                    "Mitsubishi", "Fiat", "Suzuki", "Renault"];
    const random = Math.floor(Math.random() * makes.length);
    return makes[random]
  }

  getRandomVIN() {
    const letters = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", 
                      "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", 
                      "W", "X", "Y", "Z"]
    const numbers = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    var randomLetters = Math.floor(Math.random() * letters.length);
    var firstLetter = letters[randomLetters]
    randomLetters = Math.floor(Math.random() * letters.length);
    var secondLetter = letters[randomLetters]
    randomLetters = Math.floor(Math.random() * letters.length);
    var thirdLetter = letters[randomLetters]
    randomLetters = Math.floor(Math.random() * letters.length);
    var fourthLetter = letters[randomLetters]
    randomLetters = Math.floor(Math.random() * letters.length);
    var fifthLetter = letters[randomLetters]
    var randomNumbers = Math.floor(Math.random() * numbers.length);
    var firstNumber = numbers[randomNumbers]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var secondNumber = numbers[randomNumbers]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var thirdNumber = numbers[randomNumbers]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var fourthNumber = numbers[randomNumbers]
    randomLetters = Math.floor(Math.random() * letters.length);
    var sixthLetter = letters[randomLetters]
    randomLetters = Math.floor(Math.random() * letters.length);
    var seventhLetter = letters[randomLetters]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var fifthNumber = numbers[randomNumbers]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var sixthNumber = numbers[randomNumbers]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var seventhNumber = numbers[randomNumbers]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var eighthNumber = numbers[randomNumbers]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var ninthNumber = numbers[randomNumbers]
    randomNumbers = Math.floor(Math.random() * numbers.length);
    var tenthNumber = numbers[randomNumbers]
    return firstLetter + secondLetter + thirdLetter + fourthLetter + fifthLetter +
            firstNumber + secondNumber + thirdNumber + fourthNumber +
            sixthLetter + seventhLetter + fifthNumber + sixthNumber +
            seventhNumber + eighthNumber + ninthNumber + tenthNumber
  }

  getRandomState(){
    const states = ["AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL",
                    "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME",
                    "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH",
                    "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "PR",
                    "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "VI", "WA",
                    "WV", "WI", "WY"];
    const random = Math.floor(Math.random() * states.length);
    return states[random]
  }
}
