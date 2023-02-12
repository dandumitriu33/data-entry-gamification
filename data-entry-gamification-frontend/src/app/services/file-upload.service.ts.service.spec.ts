import { TestBed } from '@angular/core/testing';

import { FileUploadServiceTsService } from './file-upload.service.ts.service';

describe('FileUploadServiceTsService', () => {
  let service: FileUploadServiceTsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FileUploadServiceTsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
