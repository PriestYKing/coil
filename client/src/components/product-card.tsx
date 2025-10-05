import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Product } from "@/types/product-type";
import { Link } from "@radix-ui/react-navigation-menu";
import Image from "next/image";
import { Badge } from "./ui/badge";

export function ProductCard({ product }: { product: Product }) {
  return (
    <div className="shadow-lg rounded-lg overflow-hidden">
      <div className="relative aspect-[2/3]">
        <Image
          src={product.picture}
          alt={product.name}
          className="object-cover hover:scale-105 transition-all duration-300"
          fill
        />
      </div>
      <div className="flex py-4 items-center flex-col">
        <h1 className="font-medium">{product.name}</h1>
        <p className="text-sm text-gray-500 ">{product.description}</p>
      </div>
      <div className="flex pb-4 items-center justify-around">
        <div>
          <Label>Stock</Label>
          <Badge variant="default">{product.stock}</Badge>
        </div>
        <div>
          <Label>Price</Label>
          <Badge variant="default">$ {product.price}</Badge>
        </div>
      </div>
    </div>
  );
}
