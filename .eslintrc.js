module.exports = {
	'env': {
		'es6': true,
		'node': true,
		"jest": true
	},
	'extends': ['eslint:recommended', 'airbnb-base/legacy'],
	'globals': {
		'Atomics': 'readonly',
		'SharedArrayBuffer': 'readonly'
	},
	'parserOptions': {
		'ecmaVersion': 2018,
		'sourceType': 'module'
	},
	'rules': {
		'indent': [
			'error',
			'tab'
		],
		'eqeqeq': [
			'error',
			'always'
		],
		'no-trailing-spaces': 'error',
		'object-curly-spacing': [
			'error', 'always'
		],
		'arrow-spacing': [
			'error', { 'before': true, 'after': true }
		],		
		'quotes': [
			'error',
			'single'
		],
		'semi': [
			'error',
			'always'
		],
		'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
		'linebreak-style': 0,
		'no-tabs': ['error', { allowIndentationTabs: true }],
		'camelcase': 'error'
	},
}