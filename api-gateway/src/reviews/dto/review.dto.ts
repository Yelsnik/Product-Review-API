import {
  IsBoolean,
  IsLowercase,
  IsNumber,
  IsOptional,
  IsString,
} from 'class-validator';

export class addReviewsDTO {
  @IsString()
  review: string;
}

export class addReviewParamsDTO {
  @IsString()
  id: string;
}

export class getProductsDTO {
  @IsNumber()
  page: number;

  @IsString()
  country: string;
}

export class getProductDetailsDTO {
  @IsString()
  productId: string;

  @IsString()
  country: string;
}
