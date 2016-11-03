function jqready() {

	var selectedItem, selectedPrice; //selected product in Order
	
	var amountDue = []
	
	var itemsPurchased = []
	
	var objDOM = function (id) {
		return document.getElementById(id);
	}	
    
	function round(value, decimals) {
		return parseFloat(Math.round(value * 100) / 100).toFixed(2)
	}
	
//http://stackoverflow.com/questions/105034/create-guid-uuid-in-javascript	
function uuid()
{
   var chars = '0123456789abcdef'.split('');

   var uuid = [], rnd = Math.random, r;
   uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
   uuid[14] = '4'; // version 4

   for (var i = 0; i < 36; i++)
   {
      if (!uuid[i])
      {
         r = 0 | rnd()*16;

         uuid[i] = chars[(i == 19) ? (r & 0x3) | 0x8 : r & 0xf];
      }
   }

   return uuid.join('')
}	

$('#edtProduct').typeahead({
	//http://tatiyants.com/how-to-use-json-objects-with-twitter-bootstrap-typeahead/
    source: function (query, process) {
        items = []
        map = {}
		
		

        //TODO: get data from server
		
		/*
        var data = [
        {"id": "1", "price": "2.50", "name": "Alaxan"},
        {"id": "2", "price": "2.00", "name": "Biogesic"},
        {"id": "3", "price": "1.75", "name": "Enfagrow 1-3yrs"},
        {"id": "4", "price": "1.00", "name": "Enfagrow 4-6yrs"},
        {"id": "5", "price": "1.25", "name": "Paracetamol"}
        ];
		
        $.each(data, function (i, state) {
            map[state.name] = state;
            items.push(state.name);
        });

        process(items);  
		*/	

/*
[{
		"Id" : 1,
		"Price" : 55,
		"Name" : "0.9% Sodium Chloride - 1L Bottle"
	}, {
		"Id" : 2,
		"Price" : 17,
		"Name" : "0.9% Sodium Chloride - 50 mL Bottle"
	}, {
		"Id" : 3,
		"Price" : 35,
		"Name" : "0.9% Sodium Chloride - 500 mL Bottle"
	}, {
		"Id" : 4,
		"Price" : 13.78,
		"Name" : "Amoxicillin - 100 mg/mL, 10 mL Drops"
	}, {
		"Id" : 5,
		"Price" : 14.25,
		"Name" : "Amoxicillin - 125 mg/5 mL, 60 mL Suspension"
	}, {
		"Id" : 6,
		"Price" : 0.76,
		"Name" : "Amoxicillin - 250 mg Capsule"
	}, {
		"Id" : 7,
		"Price" : 19.13,
		"Name" : "Ascorbic Acid (Vitamin C) - 100 mg/5 mL, 120 mL Syrup"
	}, {
		"Id" : 8,
		"Price" : 12.6,
		"Name" : "Ascorbic Acid (Vitamin C) - 100 mg/mL, 15 mL Drops"
	}, {
		"Id" : 9,
		"Price" : 124.25,
		"Name" : "Erythromycin - 0.5%, 3.5g Eye Ointment Tube"
	}, {
		"Id" : 10,
		"Price" : 11.33,
		"Name" : "Glucose (Dextrose) - 50%, 50 mL Vial"
	}
]
*/		

		
		//http://stackoverflow.com/questions/12621823/ajax-call-populate-typeahead-bootstrap

        $.getJSON('/products', function (data) {
			$.each(data, function (i, state) {
				//alert(state.Name)
				map[state.Name] = state
				items.push(state.Name)
			})
        })
		//alert(items)
		setTimeout(function(){
				process(items)
		}, 700) //milliseconds
		
    },

    updater: function (item) {
        selectedItem = map[item].Id;
        selectedPrice = map[item].Price;
        return item;        
    },

	
    matcher: function (item) {
        if (item.toLowerCase().indexOf(this.query.trim().toLowerCase()) != -1) {
            return true;
        }        
    },
	

    sorter: function (items) {
        return items.sort();
    },

    highlighter: function (item) {
        var regex = new RegExp( '(' + this.query + ')', 'gi' );
        return item.replace( regex, "<strong>$1</strong>" );       
    },
});	



	$("#btnGetPayment").click(function () {		
		var rowCount = $('#tblOrderDetails tr').length - 1;
		if (rowCount <= 0) {
			return
		}
		
		$('#modalGetPayment').modal('show')
		
		$('#edtPymtAmtDue').val($('#edtAmtDue').val())
		
		$('#edtAmtTendered').focus()
	});  
	
	function RenderItemsTable() {
		$("#tblOrderDetails > tbody").html("")
		
		_.each(itemsPurchased, function(obj) {
			//console.log(obj.uuid + ' ' + obj.qty + ' ' + obj.product + ' @ ' + obj.price)
			$("#tblOrderDetails").find('tbody')
				.append($('<tr>')
					.append($('<td>')
						.text(obj.qty + ' ' + 
							obj.product + ' @ ' +
							obj.price)
					)
                .append($('<td>')
                    .text(obj.total)
                )                
            )			
		})		
	}
	
	function AddItems() {
		if ($("#edtProduct").val() == "" || $("#edtQty").val() == "") {
			alert("Please specify product and quantity")
			return
		}

		//compute itemTotal
		var i = $('#edtPrice').val() * $('#edtQty').val()
		//http://stackoverflow.com/questions/6134039/format-number-to-always-show-2-decimal-places
		var itemTotal = round(i, 2)		
		$('#edtItemTotal').val(itemTotal)
		
        $("#tblOrderDetails").find('tbody')
            .append($('<tr>')
                .append($('<td>')
                    .text($('#edtQty').val() + ' ' + 
                        $('#edtProduct').val() + ' @ ' +
                        $('#edtPrice').val())
                )
                .append($('<td>')
                    .text($('#edtItemTotal').val())
                )                
            )
		
		var i = parseFloat($('#edtItemTotal').val())
		amountDue.push(i)
		
		//this is the model part of order details
		itemsPurchased.push({
			uuid: uuid(),
			qty: $('#edtQty').val(),
			product: $('#edtProduct').val(),
			price: $('#edtPrice').val(),
			total: $('#edtItemTotal').val()
		})
		
		//alert(_.values(itemsPurchased[0]))
		//alert(itemsPurchased.length)
		
		var v = round(_.sum(amountDue), 2)
		//alert(v)
		$('#edtAmtDue').val(v)				
			
		$('#edtProduct').val('')
		$('#edtQty').val('')
		$('#edtPrice').val('')
		$('#edtItemTotal').val('')				
	}
   
    $('#edtProduct').change(function() {
		
        $('#edtPrice').val(selectedPrice)
    });
	
	function RenderAmountDue() {
		var v = round(_.sum(amountDue), 2)
		//alert(v)
		$('#edtAmtDue').val(v)			
	}

/*
    $('#btnGetPayment').click(function() {
		//$('#modalUpdateDeleteItems').modal('show')
		var rowCount = $('#tblOrderDetails tr').length - 1;
		//alert(rowCount)
		$('#edtPymtAmtDue').val($('#edtAmtDue').val())
    });
*/
	$('#edtQty').on("keyup", function() {
		if (isNaN($(this).val()) ) {
			return
		}
		
		//compute itemTotal
		var i = $('#edtPrice').val() * $('#edtQty').val()
		//http://stackoverflow.com/questions/6134039/format-number-to-always-show-2-decimal-places
		var itemTotal = round(i, 2)		
		$('#edtItemTotal').val(itemTotal)		
	})
	
	//http://stackoverflow.com/questions/979662/how-to-detect-pressing-enter-on-keyboard-using-jquery
	$( "#edtQty" ).on( "keydown", function(event) {
		var thisVal = $(this).val()
		
		if (isNaN(thisVal) ) {
			return
		}

		if (thisVal < 1) {
			return
		}
		if(event.which == 13) {
			//alert("Entered!");
			AddItems()
			
			$('#edtProduct').focus()
		}
    });	
	
	function IsAmountTenderedInvalid() {
		var thisVal = $('#edtAmtTendered').val() * 100
		
		var due = $('#edtPymtAmtDue').val() * 100
		
		var change = round((thisVal - due)/100, 2)
		
		$('#edtAmtChange').val(change)
		
		return (change < 0.00) 
	}
	
	$('#edtAmtTendered').on("keyup", function() {
		if (isNaN($(this).val()) ) {
			return
		}

		if (IsAmountTenderedInvalid()) {
			$('#labelMsg').html('Amount tendered is insufficient')
		} else {
			$('#labelMsg').html('')
		}
	})

	$("#edtAmtTendered").on( "keydown", function(event) {
		if(event.which != 13) {
			return
		}
		
		if (IsAmountTenderedInvalid()) {
			return
		}
			
		//TODO: save to database
			
		CleanUpControls()
    });	
	
	function CleanUpControls() {
		//clean up html controls and our arrays
		$('#modalGetPayment').modal('hide');
		$('#edtProduct').focus()
		
		$('#edtPymtAmtDue').val('')
		$('#edtAmtTendered').val('')
		$('#edtAmtChange').val('')	
		$('#edtAmtDue').val('')

		//http://stackoverflow.com/questions/723112/jquery-fastest-way-to-remove-all-rows-from-a-very-large-table
		$("#tblOrderDetails > tbody").html("")
		
		$('#cmbItemsPurchased').empty()
		
			
		//http://stackoverflow.com/questions/1232040/how-do-i-empty-an-array-in-javascript
		itemsPurchased.length = 0		
		amountDue.length = 0
	}
	
	$('#btnCancelOrder').click(function() {
		CleanUpControls()
	})
	
	function renderCmbItemsPurchased() {
		$('#cmbItemsPurchased').empty()
		amountDue.length = 0
		
		$.each(itemsPurchased, function(i, el) { 
			$('#cmbItemsPurchased').append( new Option(el.qty + ' ' + el.product + ' @ ' + el.price, el.uuid) )
			
			//update amountDue
			var i = parseFloat(el.total)
			amountDue.push(i)
		})
	}
	
	$('#btnDeleteItem').click(function() {
		/*
		itemsPurchased.push({
			uuid: uuid(),
			qty: $('#edtQty').val(),
			product: $('#edtProduct').val(),
			price: $('#edtPrice').val(),
			total: $('#edtItemTotal').val()
		})
		*/
		//http://stackoverflow.com/questions/6679134/add-options-dynamically-in-combo-box-using-jquery
		
		//$('#cmbItemsPurchased').empty()
		$('#modalDeleteItem').modal('show')
		
		//alert(_.object(itemsPurchased))

		renderCmbItemsPurchased()
	})
	
	$('#btnDeleteSelectedItem').click(function() {
		
		var select = objDOM('cmbItemsPurchased');
		//alert(select.options[select.selectedIndex].value)
		var uuid = select.options[select.selectedIndex].value
		
		var arrObj = _.reject(itemsPurchased, function(obj){ 
			return obj.uuid === uuid
		})
		
		//side effects
		itemsPurchased.length = 0
		itemsPurchased = arrObj.slice(0)		
		renderCmbItemsPurchased()
		RenderItemsTable()
		RenderAmountDue()
		
		//WAIT
		
		//alert(_.values(_.first(arrObj)))
		/*
		_.each(arrObj, function(obj) {
			console.log(obj.uuid + ' ' + obj.qty + ' ' + obj.product + ' @ ' + obj.price)
		})
		*/
		
	})
}
