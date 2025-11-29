import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    MongooseModule.forRoot(process.env.MONGO_URI!), //Não é o ideal, mas em um ambiente onde eu sei que o .env sempre terá a variável, posso usar o "!" para garantir que não é nulo
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
