import {
  Body,
  Controller,
  Get,
  Param,
  Post,
  Query,
  Res,
  UseFilters,
} from '@nestjs/common';
import { Response } from 'express';
import { ReviewsService } from './reviews.service';
import {
  addReviewParamsDTO,
  addReviewsDTO,
  getProductDetailsDTO,
  getProductsDTO,
} from './dto/review.dto';
import { AddReviewRequest, GetReviewsRequest } from 'pb/review_service';
import { RpcToHttpExceptionFilter } from 'src/exceptions/http-exception.filter';
import {
  GetProductDetailsRequest,
  GetProductsRequest,
} from 'pb/product_service';

@Controller({ path: 'reviews', version: '1' })
@UseFilters(RpcToHttpExceptionFilter)
export class ReviewsController {
  constructor(private readonly reviewsService: ReviewsService) {}

  @Post(':id')
  async addReview(
    @Param() params: addReviewParamsDTO,
    @Body() body: addReviewsDTO,
    @Res() response: Response,
  ) {
    const request: AddReviewRequest = {
      review: body.review,
      productId: params.id,
    };
    console.log(params.id);
    const review = await this.reviewsService.addReview(request);

    response.status(201).json({
      message: 'success',
      data: review,
    });
  }

  @Get(':id')
  async getReviews(@Param() params: any, @Res() response: Response) {
    const request: GetReviewsRequest = {
      productId: params.productId,
    };

    const reviews = await this.reviewsService.getReviews(request);

    response.status(200).json({
      message: 'success',
      data: reviews,
    });
  }

  @Get('products')
  async getProducts(@Query() query: getProductsDTO, @Res() response: Response) {
    console.log('1', query);
    const country = query.country.toUpperCase();
    const request: GetProductsRequest = {
      page: query.page.toString(),
      country: country,
    };

    console.log(request);

    const products = await this.reviewsService.getProducts(request);

    response.status(200).json({
      message: 'success',
      data: products,
    });
  }

  @Get('product-details')
  async getProductDetails(
    @Query() query: getProductDetailsDTO,
    @Res() response: Response,
  ) {
    const request: GetProductDetailsRequest = {
      asin: query.productId,
      country: query.country,
    };

    const product = await this.reviewsService.getProductDetails(request);

    response.status(200).json({
      message: 'success',
      data: product,
    });
  }
}
