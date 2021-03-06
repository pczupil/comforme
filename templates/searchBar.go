package templates

const SearchBar = `
	<form method="get" action="/search">
		<div class="row collapse">
			<div class="small-10 columns">
			<input type="text" placeholder="Page Search" name="q" id="page-search-textbox" />
			</div>
			<div class="small-2 columns">
				<button type="submit" class="button postfix">Search</button>
			</div>
		</div>
	</form>

	<script src="//cdn.jsdelivr.net/algoliasearch/3/algoliasearch.min.js"></script>
	<script src="//cdn.jsdelivr.net/autocomplete.js/0/autocomplete.min.js"></script>
	<script>
		var client = algoliasearch({{ .appId }}, {{ .publicSearchKey }})
		var index = client.initIndex('Pages');
		autocomplete('#page-search-textbox', {hint: false}, [
			{
				source: autocomplete.sources.hits(index, {hitsPerPage: 5}),
				displayKey: 'title',
				templates: {
					suggestion: function(suggestion) {
						return suggestion._highlightResult.title.value;
					}
				}
			}
		]).on('autocomplete:selected', function(event, suggestion, dataset) {
			console.log(suggestion, dataset);
		});
	</script>

`
