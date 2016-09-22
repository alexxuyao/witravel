$(function() {

    var travellist = Vue.extend({
        template : $('#travellist_view').html()
    });
    
    var tourguide = Vue.extend({
        template : '<p>This is tourguide!</p>'
    });

    var app = Vue.extend({});

    var router = new VueRouter();

    router.map({
        '/travellist' : {
            component : travellist
        },
        '/tourguide' : {
            component : tourguide
        },
    });

    router.start(app, '#app');
});