<script>
	import '../app.postcss';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });
</script>

<!-- App Shell -->
<!-- svelte-ignore a11y-invalid-attribute -->
<AppShell>
	<svelte:fragment slot="header">
		<!-- App Bar -->
		<AppBar background="nav-bar">
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase"> <a href="/" class="logo-text">Adapt Ecomm</a></strong>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<a class="btn btn-sm nav-link" href="/login"> Login </a>
				<a class="btn btn-sm nav-link" href="#"> Signup </a>
			</svelte:fragment>
			<svelte:fragment slot="default">
				<nav class="categories ml-6">
					<a class="category-link nav-link" href="/electronics">Electronics</a>
					<a class="category-link nav-link" href="/fashion">Fashion</a>
					<a class="category-link nav-link" href="/home-appliances">Home Appliances</a>
					<a class="category-link nav-link" href="/books">Books</a>
					<a class="category-link nav-link" href="/toys">Toys</a>
				</nav>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<!-- Page Route Content -->
	<slot />
	<!-- Footer -->
	<svelte:fragment slot="pageFooter">
		<footer class="text-center py-4 page-footer">
			&copy; 2021 Adapt Ecomm. All rights reserved.
		</footer>
	</svelte:fragment>
</AppShell>

<style>
	.category-link {
		margin-left: 10px;
		margin-right: 10px;
	}
</style>
