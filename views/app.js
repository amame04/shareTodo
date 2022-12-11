var app = new Vue({
  el: "#app",
  data: {
  },
  methods:{
    post: function(){
      config = {
        headers:{
          'X-Requested-With' : 'XMLHttpRequest',
          'Content-Type':'application / x-www-form-urlencoded'
        },
        withCredentials:true,
      }
      param = JSON.parse("{\"id\" : \"test\"}")
      axios.post("/",param,config)
      .then(function(res){
        app.result = res.data
      })
      .catch(function(res){
        app.result = res.data
      })
    }
  }
})