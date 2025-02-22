import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { GetProductDetailsRequest, GetProductsRequest, ProductClient } from 'pb/product_service';
import {
  AddReviewRequest,
  GetReviewsRequest,
  ReviewClient,
} from 'pb/review_service';
import { last, lastValueFrom } from 'rxjs';

@Injectable()
export class ReviewsService implements OnModuleInit {
  reviewService: ReviewClient;
  productService: ProductClient;

  constructor(
    @Inject('REVIEW_SERVICE') private reviewClient: ClientGrpc,
    @Inject('PRODUCT_SERVICE') private productClient: ClientGrpc,
  ) {}

  onModuleInit() {
    this.reviewService = this.reviewClient.getService<ReviewClient>('Review');
    this.productService =
      this.productClient.getService<ProductClient>('Product');
  }

  async addReview(request: AddReviewRequest) {
    const result = this.reviewService.addReview(request);

    const review = await lastValueFrom(result);

    return review;
  }

  async getReviews(request: GetReviewsRequest) {
    const result = this.reviewService.getReviews(request);

    const reviews = await lastValueFrom(result);

    return reviews.reviews;
  }

  async getProducts(request: GetProductsRequest) {
    const result = this.productService.getProducts(request)

    const products = await lastValueFrom(result)

    return products.product
  }

  async getProductDetails(request: GetProductDetailsRequest){
    const result = this.productService.getProductDetails(request)

    const product = await lastValueFrom(result)

    return product.product
  }
}
