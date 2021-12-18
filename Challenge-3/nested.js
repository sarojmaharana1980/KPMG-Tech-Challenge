var object, nested;

//object = {“a”: {“b”: {“c”: ”d”}}};
//key = "a/b/c"



nested = function(object, key) {
	  var pairs;
	  //get pair length
	  if (!(pairs = _(object).pairs()).length) {
	    return [
	      {
	        keys: key,
	        value: object
	      }
	    ];
		console.log(object)
	  } else {
	    return [].concat.apply([], _(pairs).map(function(kv) {
	      return nested(kv[1], key.concat(kv[0]));
	    }));
	  }
	};