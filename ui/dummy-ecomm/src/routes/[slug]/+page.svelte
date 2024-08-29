<script>
	let maxPriceFromRange = 0;
	let products = [
		{
			thumbnail: 'https://placehold.co/400x400',
			title: 'Mawa Smartphone Pro Max Ultra',
			rating: 4.5,
			description: [
				'8GB DDR7 Ram',
				'50MP Front Camera, 64MP Rear Camera',
				'128 GB Internal Storage'
			],
			price: 49999,
			mrp: 60000,
			id: 1
		},
		{
			thumbnail: 'https://placehold.co/400x400',
			title: 'Mango Fold Pro',
			rating: 4.1,
			description: ['4GB Ram', '16MP Front Camera, 50MP Rear Camera', '64 GB Internal Storage'],
			price: 19999,
			mrp: 35000,
			id: 2
		},
		{
			thumbnail: 'https://placehold.co/400x400',
			title: 'OneMinus 20T',
			rating: 4.3,
			description: ['6GB Ram', '32MP Front Camera, 48MP Rear Camera', '128 GB Internal Storage'],
			price: 29999,
			mrp: 40000,
			id: 3
		}
		// Add more products as needed
	];
</script>

<div class="container">
	<div class="filters">
		<h2>Filters</h2>
		<hr />
		<div class="filter">
			<label for="priceRange">Price</label>
			<input
				type="range"
				id="priceRange"
				name="priceRange"
				min="0"
				max="99999"
				step="100"
				bind:value={maxPriceFromRange}
			/>
			<small>
				Max Price: {maxPriceFromRange}
			</small>
		</div>
		<div class="filter">
			<h3>Brands</h3>
			<div class="flex items-center space-x-1">
				<input type="checkbox" id="brand1" name="brand1" value="brand1" />
				<label for="brand1">Mawa</label>
			</div>
			<div class="flex items-center space-x-2">
				<input type="checkbox" id="brand2" name="brand2" value="brand2" />
				<label for="brand2">Mango</label>
			</div>
			<div class="flex items-center space-x-2">
				<input type="checkbox" id="brand3" name="brand3" value="brand3" />
				<label for="brand3">Maomi</label>
			</div>
			<div class="flex items-center space-x-2">
				<input type="checkbox" id="brand4" name="brand4" value="brand4" />
				<label for="brand4">OneMinus</label>
			</div>
		</div>
		<div class="filter">
			<h3>Customer Rating</h3>
			<div class="flex items-center space-x-1">
				<input type="checkbox" id="rating4up" name="rating4up" value="4" />
				<label for="rating4up">4<span class="icon">★</span> and up</label>
			</div>
			<div class="flex items-center space-x-2">
				<input type="checkbox" id="rating3up" name="rating3up" value="3" />
				<label for="rating3up">3<span class="icon">★</span> and up</label>
			</div>
		</div>
	</div>
	<div class="product-listing">
		{#each products as product}
			<div class="product">
				<div class="product-thumbnail">
					<img class="thumbnail" src={product.thumbnail} alt={product.title} />
				</div>
				<div class="product-details">
					<h3 class="title space-y-5"><a href={'/product/' + product.id}>{product.title}</a></h3>
					<p class="rating">{product.rating} <span class="icon">★</span></p>
					<ul class="desc">
						{#each product.description as desc}
							<li>{desc}</li>
						{/each}
					</ul>
				</div>
				<div class="product-price">
					<p class="sale-price">INR {product.price}</p>
					<p class="mrp"><strike>MRP: {product.mrp}</strike></p>
				</div>
			</div>
		{/each}
	</div>
</div>

<style>
	.container {
		display: flex;
	}
	.filters {
		width: 20%;
		padding: 10px;
		border-right: 1px solid #ccc;
	}
	.product-listing {
		width: 80%;
		padding: 10px;
	}
	.product {
		display: flex;
		border: 1px solid #ccc;
		padding: 10px;
		margin-bottom: 10px;
	}
	.product-thumbnail {
		width: 20%;
	}
	.product-details {
		width: 60%;
		padding-left: 10px;
	}
</style>
