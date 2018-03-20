import { Component, EventEmitter, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-warning',
  templateUrl: './warning.component.html',
  styleUrls: ['./warning.component.scss']
})
export class WarningComponent implements OnInit {

  @Output() onSubmit = new EventEmitter();

  constructor() { }

  ngOnInit() {
  }

  next() {
    this.onSubmit.emit();
  }
}
