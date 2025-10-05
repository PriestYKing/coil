import { AppSidebar } from "@/components/app-sidebar";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Frown, Plus } from "lucide-react";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Checkbox } from "@/components/ui/checkbox";
import { ProductCard } from "@/components/product-card";
import { Product } from "@/types/product-type";

const products: Product[] = [
  {
    id: 1,
    name: "Voldemort",
    description: "A book of spells",
    price: 23,
    sizes: ["S", "M", "L"],
    picture: "/1g.png",
    stock: 10,
  },
  {
    id: 2,
    name: "Harry Potter",
    description: "A book of spells",
    price: 23,
    sizes: ["S", "M", "L"],
    picture: "/1gr.png",
    stock: 8,
  },
  {
    id: 3,
    name: "Dumbledore",
    description: "A book of spells",
    price: 23,
    sizes: ["S", "M", "L"],
    picture: "/6g.png",
    stock: 5,
  },
  {
    id: 4,
    name: "Harry Potter",
    description: "A book of spells",
    price: 23,
    sizes: ["S", "M", "L"],
    picture: "/1gr.png",
    stock: 8,
  },
  {
    id: 5,
    name: "Dumbledore",
    description: "A book of spells",
    price: 23,
    sizes: ["S", "M", "L"],
    picture: "/6g.png",
    stock: 5,
  },
];

export default function Page() {
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <header className="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12">
          <div className="flex items-center gap-2 px-4">
            <SidebarTrigger className="-ml-1" />
            <Separator
              orientation="vertical"
              className="mr-2 data-[orientation=vertical]:h-4"
            />
            <Breadcrumb>
              <BreadcrumbList>
                <BreadcrumbItem className="hidden md:block">
                  <BreadcrumbLink href="#">Inventory</BreadcrumbLink>
                </BreadcrumbItem>
                <BreadcrumbSeparator className="hidden md:block" />
                <BreadcrumbItem>
                  <BreadcrumbPage>Dashboard</BreadcrumbPage>
                </BreadcrumbItem>
              </BreadcrumbList>
            </Breadcrumb>
          </div>
        </header>
        <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
          {products.length == 0 ? (
            <div className="flex items-center justify-center h-screen">
              <Sheet>
                <SheetTrigger asChild>
                  <Button>
                    <Plus />
                    Add Product
                  </Button>
                </SheetTrigger>
                <SheetContent>
                  <SheetHeader>
                    <SheetTitle>Add Product</SheetTitle>
                    <SheetDescription>
                      Enter details for your product.
                    </SheetDescription>
                  </SheetHeader>
                  <div className="grid flex-1 auto-rows-min gap-6 px-4">
                    <div className="grid gap-3">
                      <Label htmlFor="sheet-demo-name">Name</Label>
                      <Input id="sheet-demo-name" placeholder="Voldemort" />
                    </div>
                    <div className="grid gap-3">
                      <Label htmlFor="sheet-demo-username">Description</Label>
                      <Input
                        id="sheet-demo-username"
                        placeholder="A book of spells"
                      />
                    </div>
                    <div className="grid gap-3">
                      <Label htmlFor="sheet-demo-username">Price</Label>
                      <Input
                        id="sheet-demo-username"
                        placeholder="$23"
                        type="number"
                      />
                    </div>
                    <Label>Sizes</Label>
                    <div className="flex items-center justify-between">
                      <div className="flex gap-2">
                        <Checkbox id="s" />
                        <Label htmlFor="s">S</Label>
                      </div>
                      <div className="flex gap-2">
                        <Checkbox id="m" />
                        <Label htmlFor="m">M</Label>
                      </div>
                      <div className="flex gap-2">
                        <Checkbox id="l" />
                        <Label htmlFor="L">L</Label>
                      </div>
                    </div>
                    <div className="grid gap-3">
                      <Label htmlFor="picture">Picture</Label>
                      <Input id="picture" type="file" />
                    </div>
                  </div>
                  <SheetFooter>
                    <Button type="submit">Save changes</Button>
                    <SheetClose asChild>
                      <Button variant="outline">Close</Button>
                    </SheetClose>
                  </SheetFooter>
                </SheetContent>
              </Sheet>
            </div>
          ) : (
            <div className="mt-2 grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-12">
              {products.map((product) => (
                <ProductCard key={product.id} product={product} />
              ))}
            </div>
          )}
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}
