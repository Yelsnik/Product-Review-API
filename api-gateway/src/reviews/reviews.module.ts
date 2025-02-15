import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { REVIEW_PACKAGE_NAME as Review } from 'pb/review_service';
import { ReviewsService } from './reviews.service';
import { ReviewsController } from './reviews.controller';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'REVIEW_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: Review,
          protoPath: join(__dirname, '../../protos/review_service.proto'),
          url: '0.0.0.0:9090',
        },
      },
      {
        name: 'PRODUCT_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: Review,
          protoPath: join(__dirname, '../../protos/product_service.proto'),
          url: '0.0.0.0:9090',
        },
      },
    ]),
  ],
  providers: [ReviewsService],
  controllers: [ReviewsController],
})
export class ReviewsModule {}
