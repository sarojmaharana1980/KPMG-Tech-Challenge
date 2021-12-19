var object, nested;

//object = {“a”: {“b”: {“c”: ”d”}}};
//key = "a/b/c"



nested = function(object, key) {
	  var mappairs;
	  //get pair length and validate
	  if (!(mappairs = _(object).mappairs()).length) {
	    return [
	      {
	        keys: key,
	        value: object
	      }
	    ];
		console.log(object)
	  } else {
		  // get data directly from map object
	    return [].concat.apply([], _(mappairs).map(function(kv) {
	      return nested(kv[1], key.concat(kv[0]));
	    }));
	  }
	};