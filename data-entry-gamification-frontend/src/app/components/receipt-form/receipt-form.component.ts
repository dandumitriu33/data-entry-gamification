import { Component } from '@angular/core';
import { Receipt } from 'src/app/entities/receipt';

@Component({
  selector: 'app-receipt-form',
  templateUrl: './receipt-form.component.html',
  styleUrls: ['./receipt-form.component.css']
})
export class ReceiptFormComponent {

  receipt: Receipt = {id: 0, modelYear: 0, make: "", vin: "", firstName: "", lastName: "", state: ""};

  onSubmitTemplateBased(receipt: Receipt) { 
    console.log("receipt: ", receipt)
  }

  newReceipt() {
    console.log("new receipt clicked")
  }

}
