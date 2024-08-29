<script>
	import { Splide, SplideSlide } from '@splidejs/svelte-splide';
	import '@splidejs/splide/dist/css/splide.min.css';

	const productImages = [
		'https://placehold.co/600x400',
		'https://placehold.co/600x400',
		'https://placehold.co/600x400'
	];

	const product = {
		title: 'Product Title',
		rating: 4.5,
		price: 500,
		deliveryInfo: 'Delivery in 3-5 business days',
		description: 'This is a detailed description of the product.',
		specs: [
			{ key: 'Weight', value: '1kg' },
			{ key: 'Dimensions', value: '10x10x10 cm' }
			// Add more specs as needed
		],
		similarProducts: [
			{ name: 'Similar Product 1', image: 'https://placehold.co/400x400', price: 100 },
			{ name: 'Similar Product 2', image: 'https://placehold.co/400x400', price: 200 }
			// Add more similar products as needed
		],
		reviews: [
			{ customer: 'John Doe', rating: 5, comment: 'Great product!' },
			{ customer: 'Jane Smith', rating: 4, comment: 'Very good, but could be improved.' }
			// Add more reviews as needed
		]
	};
</script>

<div class="space-y-8 p-8">
	<!-- First Section: Product Gallery and Info -->
	<section class="flex">
		<div class="product-gallery">
			<Splide options={{ type: 'loop', perPage: 1, pagination: true, arrows: true }}>
				{#each productImages as image}
					<SplideSlide>
						<!-- svelte-ignore a11y-img-redundant-alt -->
						<img src={image} alt="Product Image" class="w-full h-auto" />
					</SplideSlide>
				{/each}
			</Splide>
		</div>
		<div class="w-1/2 space-y-4 p-8">
			<h1 class="text-3xl font-bold product-title">{product.title}</h1>
			<div class="flex items-center space-x-2">
				<span class="text-yellow-500">{'★'.repeat(Math.floor(product.rating))}</span>
				<span class="text-gray-500">({product.rating})</span>
			</div>
			<p class="text-gray-700">{product.deliveryInfo}</p>
			<p class="text-2xl">INR {product.price}</p>
		</div>
	</section>

	<!-- Second Section: Product Details -->
	<section class="space-y-4 p-8">
		<h2 class="text-2xl font-bold">Product Details</h2>
		<p class="text-gray-700">{product.description}</p>
		<table class="min-w-full bg-white">
			<tbody>
				{#each product.specs as spec}
					<tr class="border-b">
						<td class="py-2 px-4 font-medium text-gray-700">{spec.key}</td>
						<td class="py-2 px-4 text-gray-700">{spec.value}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</section>

	<!-- Third Section: Similar Products -->
	<section class="space-y-4 p-8">
		<h2 class="text-2xl font-bold">Similar Products</h2>
		<div class="flex space-x-4 overflow-x-auto">
			{#each product.similarProducts as similarProduct}
				<div class="min-w-[200px] bg-white p-4 rounded-lg shadow-sm flex-shrink-0 border-2">
					<img
						src={similarProduct.image}
						alt={similarProduct.name}
						class="w-full h-48 object-cover"
					/>
					<h3 class="text-lg font-semibold mt-2">{similarProduct.name}</h3>
					<p class="text-gray-500 mt-1">₹{similarProduct.price}</p>
				</div>
			{/each}
		</div>
	</section>

	<!-- Fourth Section: Customer Reviews -->
	<section class="space-y-4 p-8">
		<h2 class="text-2xl font-bold">Customer Reviews</h2>
		{#each product.reviews as review}
			<div class="bg-white p-4 rounded-lg shadow-md">
				<div class="flex items-center space-x-2">
					<span class="text-yellow-500">{'★'.repeat(review.rating)}</span>
					<span class="text-gray-500">({review.rating})</span>
				</div>
				<p class="text-gray-700 mt-2">{review.comment}</p>
				<p class="text-gray-500 text-sm mt-1">- {review.customer}</p>
			</div>
		{/each}
	</section>
</div>

<style>
	.product-gallery {
		width: 40%;
	}
	.product-title {
		color: #333;
		font-size: 2rem;
	}
</style>
