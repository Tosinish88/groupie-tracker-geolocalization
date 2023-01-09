
// filtering image on main page

function filterImage() {
  let input, filter, card, crd, cd, i, txtValue;
  input = document.getElementById("myInput");
  filter = input.value.toUpperCase();
  card = document.getElementById("cards");
  crd = card.getElementsByTagName("div");
  const x = document.getElementById("character");
  name = document.getElementsByClassName("Name");
  if (input.value=="") {
      x.innerHTML = "";
      return
  }
  //  populating the character div with filtered data
  x.innerHTML = xdata;
  // Loop through all image rows, and hide those who don't match the search query
  for (i = 0; i < crd.length; i++) {
      cd = crd[i].getElementsByTagName("div")[0];
      if (cd) {
          txtValue = cd.textContent || cd.innerText;
          if (txtValue.toUpperCase().indexOf(filter) > -1) {
              crd[i].style.display = "";
          } else {
          crd[i].style.display = "none";
          }
      }
  }
}

function completeDataList(e) {
  const fill = val => document.getElementById('myInput').innerHTML = val;
  if(!e.target.value) {
      fill(val.value.reduce((sum, [html]) => sum + html, ''));
  } else if(!(e instanceof InputEvent)) { // OR: else if(!e.inputType)
      e.target.blur();
  } else {
      const inputValue = e.target.value.toLowerCase();
      let result = '';
      for (const [html, valuePattern] of val.value) {
          if (!valuePattern.indexOf(inputValue)) {
              result += html;
          } else if (result) {
              break;
          }
      }
      fill(result);
  }
}

// for map section on artist page
let locarr = [];
    let locarr2 = [];
    // async function to wait for api request promise to resolve
    async function geocode() {
      
      let loc = document.getElementsByClassName("placemark");
      for (let i=0; i<loc.length; i++) {
        city =loc[i].innerHTML;
            await axios.get('https://maps.googleapis.com/maps/api/geocode/json?', {
        params: {
          address: city,
          key: 'Add API KEY'
        }
      })
      .then(function(response){
        console.log(response);
        //get lng and lat
        let lat = response.data.results[0].geometry.location.lat;
        let lng = response.data.results[0].geometry.location.lng;
        let location = response.data.results[0].formatted_address;
        locarr.push({lat: lat, lng: lng});
        locarr2.push(location);
      })
      .catch(function(error){
        console.log(error);
      })
      }
      console.log(locarr, locarr2);
      return locarr;
    }
    // google default map initialization
    function initMap() {
        // geocoding
      geocode().then(function(locarr) {
        let map = new google.maps.Map(document.getElementById("map"), {
          zoom: 2,
          center: { lat: 0, lng: 0},
        });
        // supplying coordinates to map
        for (let i=0; i<locarr.length; i++) {
          let marker = new google.maps.Marker({
            position: locarr[i],
            map: map,
          });
          // adding info window to map
          let infowindow = new google.maps.InfoWindow({
            content: locarr2[i],
          });
          // adding event listener to info window so it respond on click
          marker.addListener("click", () => {
            infowindow.open(map, marker);
          });
        }
      })
  }
  // calling initMap function
  window.initMap = initMap;
