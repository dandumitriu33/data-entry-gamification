import { EventEmitter } from '@angular/core';

export class Emitters {
    static authEmitter = new EventEmitter<boolean>();
    static inputEmitter = new EventEmitter<boolean>();
    static uploadAvatarEmitter = new EventEmitter<boolean>();
}