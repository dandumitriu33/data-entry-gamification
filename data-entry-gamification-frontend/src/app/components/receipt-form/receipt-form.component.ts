import { Component } from '@angular/core';
import { Receipt } from 'src/app/entities/receipt';

@Component({
  selector: 'app-receipt-form',
  templateUrl: './receipt-form.component.html',
  styleUrls: ['./receipt-form.component.css']
})
export class ReceiptFormComponent {

  model = new Receipt(0, 1999, "Honda", "ABC123", "John", "Johnson", "NY")
  submitted = false
  onSubmit() { this.submitted = true}

}
