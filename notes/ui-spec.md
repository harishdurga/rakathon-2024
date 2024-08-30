# UI Spec
## Customer Segments
### Age Groups
#### Young Aged
Age is between 18-24 years

#### Middle Aged
Age is between 35-50 years

#### Old Aged
Age is above 51+ years

### Gender
#### Male
#### Female

### Location
A state in India

### Purchase Value
#### High Value Customer
Checks products with high price range, brand conscious, and prefers quality over price.
#### Low Value Customer
Checks products with low price range, prefers discounts and offers, and price sensitive.

## Different Types of Sections/Components
### Banner Slider
### Type: 
main-banner-slider 
### Description:
This is a banner slider that will be used to display the different banners on the website.
### filters 
these are the filters that will be used to filter the banners that will be displayed within this component.
- **categories:** comma separated list of categories that the banners belong to. Ex: 1,2

### Category Products
### Type:
category-products
### Description:
In this section, we will display the products that belong to a specific category. You can add multiple category products sections in the layout to display products from different categories by applying different filters based on the prompt and requirements.
#### NOTE: You can add upto 4 category products sections in the layout.
### Display Title:
This title will be displayed above the products that belong to a specific category. You can change the display title as per the age group in the prompt. Be creative with the title by including the name of the Category.  Don't be too long with the title. Keep it compact.
- if the audience is a young and the category is Electronics, the title can be like "Electronics you love".
- if the audience is a middle aged person and the category is Home & Kitchen, the title can be like "Home & Kitchen essentials".
- if the audience is a young person and the category is Fashion, the title can be like "Fashion you love".

### filters
- **category_id:** the id of the category whose products will be displayed.
- **limit:** the number of products that will be displayed in this section. 
- **sort:** the order in which the products will be displayed. Ex: price_asc, price_desc, name_asc, name_desc
- **keyword**: the keyword that will be used to filter the products. Ex: "dhoti", "saree"
- **min_price**: the minimum price of the products that will be displayed. This will be slightly below the average purchase value for low value customers. For high value customers, this will be slightly below the average purchase value.
- **max_price**: the maximum price of the products that will be displayed. Slightly above the max purchase amount for low value customers. Don't include this filter for high value customers.
#### NOTE: Don't inlcude max_price filter for high value customers. Only include min_price filter for high value customers.

#### When to Use Keyword Filter
- Use the keyword filter when the prompt includes specific events, festivals, or themes that require highlighting particular products.
- Examples:
  - If the prompt mentions a festival like Christmas, use keywords like "Christmas", "gifts", "decorations".
  - If the prompt mentions traditional clothing for a festival, use keywords like "dhoti", "saree", "kurta".

## Available Product Categories
<!-- Give me table here -->
| ID | Name
|----|-----
| 1  | Electronics
| 2  | Clothing
| 3  | Home & Kitchen
| 4  | Books
| 5  | Sports & Fitness
| 6  | Beauty & Personal Care
| 7  | Toys & Games
| 8  | Automotive
| 9  | Health & Wellness
| 10 | Grocery & Gourmet Foods


## Examples
### Example 1
#### Prompt
A young male from karnataka india who likes electronics, fashion and is a high value customer.
```json
{
    "layout":[
        {
            "type": "main-banner-slider",
            "filters": {
                "categories": "1,2"
            }
        },
        {
            "type": "category-products",
            "filters": {
                "category_id": 1,
                "limit": 5,
                "sort": "price_asc",
                "min_price": 20000
            },
            "display_title": "Electronics you need"
        },
        {
            "type": "category-products",
            "filters": {
                "category_id": 2,
                "limit": 5,
                "sort": "price_asc"
            },
            "display_title": "Fashion you love"
        }
    ]
}

```
### Example 2
#### Prompt
A middle aged person, female from Andhra Pradesh india who likes Beauty & Personal Care, Home & Kitchen and is a low value customer.
```json
{
    "layout":[
        {
            "type": "main-banner-slider",
            "filters": {
                "categories": "6,3"
            }
        },
        {
            "type": "category-products",
            "filters": {
                "category_id": 6,
                "limit": 5,
                "sort": "price_asc",
                "max_price": 2000
            },
            "display_title": "Beauty Products that will make you shine"

        },
        {
            "type": "category-products",
            "filters": {
                "category_id": 3,
                "limit": 5,
                "sort": "price_asc",
                "max_price": 5000
            },
            "display_title": "Home & Kitchen essentials"
        }
    ]
}
```

### Example 3
#### Prompt
A young aged person, male from tamilnadu india who likes Fashion and Books is a low value customer. Also it's pongal season in Tamilnadu. People wear traditional clothes like dhoti and saree during pongal.
```json
{
    "layout":[
        {
            "type": "main-banner-slider",
            "filters": {
                "categories": "2,4"
            }
        },
        {
            "type": "category-products",
            "filters": {
                "category_id": 2,
                "limit": 5,
                "sort": "price_asc",
                "keyword": "dhoti",
                "max_price": 2000
            },
            "display_title": "Clothing that make feel you special"
        },
        {
            "type": "category-products",
            "filters": {
                "category_id": 4,
                "limit": 5,
                "sort": "price_asc",
                "max_price": 2000
            },
            "display_title": "Books that you can immerse yourself in"
        }
    ]
}
