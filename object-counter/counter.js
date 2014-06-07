function Counter() {
    this.elements = {};
}
Counter.prototype.lookup = function (element) {
    var el, i;
    el = this.elements[element] || [];
    for (i = 0; i < el.length; i++) {
        if (el[i][0] === element) {
            return el[i][1];
        }
    }
};
Counter.prototype.add = function (element) {
    var el, found, i;
    el = this.elements[element] || [];
    found = false;
    for (i = 0; i < el.length; i++) {
        if (el[i][0] === element) {
            el[i][1]++;
            found = true;
        }
    }
    if (!found) {
        el.push([element, 1]);
    }
    this.elements[element] = el;
};
Counter.prototype.consume = function (elements_list) {
    elements_list.map(function (elements) {
        elements.map(this.add, this);
    }, this);
};
Counter.prototype.filter = function (filter_func) {
    var result, key;
    result = [];
    for (key in this.elements) {
        if (this.elements.hasOwnProperty(key)) {
            result = result.concat(this.elements[key].filter(filter_func));
        }
    }
    return result;
};
Counter.prototype.gt = function (num) {
    return this.filter(function (tuple) {
        return tuple[1] > num;
    }).map(function (tuple) {
        return tuple[0];
    });
};
Counter.prototype.lt = function (num) {
    return this.filter(function (tuple) {
        return tuple[1] < num;
    }).map(function (tuple) {
        return tuple[0];
    });
};
Counter.prototype.eq = function (num) {
    return this.filter(function (tuple) {
        return tuple[1] == num;
    }).map(function (tuple) {
        return tuple[0];
    });
};

function uniq(array) {
    var result, i;
    result = [];
    for (i = 0; i < array.length; i++) {
        if (result.indexOf(array[i]) === -1) {
            result.push(array[i]);
        }
    }
    return result;
}

(function () {
    var sets, counter;
    sets = [
        [1, 2, 3],
        [4, 3, 5],
        [99, "asdf", 20],
        [1, 7, 3],
        [55, 55, 55, "99"],
        [98, "asdf", 21],
    ];
    counter = new Counter();
    counter.consume(sets.map(uniq));
    console.log(counter.gt(1));
})();
