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
}
