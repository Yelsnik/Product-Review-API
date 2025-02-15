import {
  IsBoolean,
  IsLowercase,
  IsNumber,
  IsOptional,
  IsString,
} from 'class-validator';

export class addReviewsDTO {
    @IsString()
    review: string
}

export class addReviewParamsDTO{
    @IsString()
    id: string
}
