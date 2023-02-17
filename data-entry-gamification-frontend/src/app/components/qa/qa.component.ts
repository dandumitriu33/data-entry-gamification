import { Component } from '@angular/core';
import { Receipt } from 'src/app/entities/receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { Emitters } from 'src/app/emitters/emitters';

@Component({
  selector: 'app-qa',
  templateUrl: './qa.component.html',
  styleUrls: ['./qa.component.css']
})
export class QaComponent {

  receipt: Receipt = {id: 0, model_year: 0, make: "", vin: "", first_name: "", last_name: "", state: "", date_added: "", qa_score: 0, qa_date: ""};

  constructor(
    private receiptService: ReceiptService
  ) {}


  onSubmitTemplateBased(receiptFromForm: Receipt) { 
    console.log("this.receipt: ", this.receipt)
    receiptFromForm.id = 0;
    console.log("receiptFromForm: ", receiptFromForm)
    // TODO: UPDATE RECEIPT
    this.receiptService.addReceipt(receiptFromForm)
      .subscribe(receiptFromForm => {
        console.log("receipt added successfully: ", receiptFromForm);
      });
    console.log("Emitting input event.");
    Emitters.inputEmitter.emit();
  }

  updateReceipt() {
    console.log("receipt update started")
  }
}
