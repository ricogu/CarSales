import { Component, OnInit } from '@angular/core';
import {OrderService} from "./order.service";
import {NotifierService} from "angular-notifier";


@Component({
  selector: 'app-order',
  templateUrl: './order.component.html',
  styleUrls: ['./order.component.css']
})



export class OrderComponent implements OnInit {
  customerName;
  batteryList;
  wheelList;
  tireList;

  showWheelSelect = false
  showTireSelect = false
  showSubmit = false
  orderSubmitted = false

  selectedBatteryId;
  selectedWheelId;
  selectedTireId;

  orderConfirmation;

  private readonly notifier: NotifierService;
  constructor(notifierService: NotifierService, private orderService : OrderService) {
    this.notifier = notifierService;
    orderService.getBatteries().subscribe(
      data => this.batteryList = data,
      error => console.error('There was an error!', error)
    )
  }

  filterWheel(batteryId: string){
    this.selectedBatteryId = batteryId
    this.orderService.getWheels(batteryId).subscribe(
      data => {console.log(data),this.wheelList = data, this.showWheelSelect = true},
      error => console.error('There was an error!', error)
    )
  }

  filterTire(wheelId: string){
    this.selectedWheelId = wheelId
    this.orderService.getTires(wheelId).subscribe(
      data => {this.tireList = data, this.showTireSelect = true},
      error => console.error('There was an error!', error)
    )
  }

  selectTire(tireId: string){
    this.selectedTireId = tireId
    this.showSubmit = true
  }

  submitOrder(){
    if(this.customerName === undefined) {
      this.notifier.notify("Warning", "name can not be empty!");
      return
    }
    let order = {
      CustomerName: this.customerName,
      BatteryId: Number(this.selectedBatteryId),
      WheelId: Number(this.selectedWheelId),
      TireId: Number(this.selectedTireId)
    }

    console.log(order)
    this.orderService.submitOrder(order).subscribe(
      data => {console.log(data), this.orderConfirmation = data , this.orderSubmitted = true},
      error => console.error('There was an error!', error)
    )
  }

  ngOnInit(): void {
  }



}


