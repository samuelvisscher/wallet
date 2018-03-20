import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { KittenTileComponent } from './kitten-tile.component';

describe('KittenTileComponent', () => {
  let component: KittenTileComponent;
  let fixture: ComponentFixture<KittenTileComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KittenTileComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KittenTileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
