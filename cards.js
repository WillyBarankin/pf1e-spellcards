window.onload = function() {
    var cards = document.getElementsByClassName("card");
    var small1 = [];
    var small2 = [];
    var small3 = [];
    for (var i = 0; i < cards.length; i++) { 
        var c = cards[i];
        var e = c.getElementsByClassName("end")[0]; 
        var cl = c.getBoundingClientRect().left; 
        var el = e.getBoundingClientRect().left;
        if (el - cl < 2) { 
            c.style.width = '63mm'; 
            c.style.columnCount = '1';
            small1.push(c);
        } else if (el - cl < 300) {
            c.style.width = '126mm';
            c.style.columnCount = '2';
            small2.push(c);
        } else if (el - cl < 500) {
            c.style.width = '189mm';
            c.style.columnCount = '3';
            small3.push(c);
        } else {
            c.style.width = '252mm';
            c.style.columnCount = '4';
        }
    }
    for (var i = 0; i < small3.length; i++) {
        document.body.appendChild(small3[i]);
        var level3 = small3[i].getElementsByClassName("number")[0].innerHTML;
        var j = 0;
        for (; j < small1.length; j++) {
            level1 = small1[j].getElementsByClassName("number")[0].innerHTML;
            if (level1 == level3) {
                break;
            }
        }
        var item = (j == small1.length) ? small1.shift() : small1.splice(j, 1)[0];
        document.body.appendChild(item);

    }
    for (var i = 0; i < small2.length; i++) {
        document.body.appendChild(small2[i]);
    }
    for (var i = 0; i < small1.length; i++) {
        document.body.appendChild(small1[i]);
    }
    void document.body.offsetHeight;
    for (var i = 0; i < cards.length; i++) {
        var c = cards[i];
        var w = c.style.width;
        if (w === '63mm') continue;
        var overflow = c.scrollHeight > c.clientHeight;
        var endEl = c.getElementsByClassName("end")[0];
        if (!overflow && endEl) {
            var cardRect = c.getBoundingClientRect();
            var endRect = endEl.getBoundingClientRect();
            overflow = endRect.bottom > cardRect.bottom + 1;
        }
        if (w === '126mm' || w === '189mm' || w === '252mm') {
            c.classList.add('tight-description');
        }
    }
    void document.body.offsetHeight;
    for (var i = 0; i < cards.length; i++) {
        var c = cards[i];
        if (!c.classList.contains('tight-description')) continue;
        var oldWidth = c.style.width;
        var e = c.getElementsByClassName("end")[0];
        var cl = c.getBoundingClientRect().left;
        var el = e.getBoundingClientRect().left;
        var delta = el - cl;
        if (delta < 2) {
            c.style.width = '63mm';
            c.style.columnCount = '1';
        } else if (delta < 300) {
            c.style.width = '126mm';
            c.style.columnCount = '2';
        } else if (delta < 500) {
            c.style.width = '189mm';
            c.style.columnCount = '3';
        } else {
            c.style.width = '252mm';
            c.style.columnCount = '4';
        }
        if (c.style.width === oldWidth) {
            c.classList.remove('tight-description');
            c.style.columnCount = oldWidth === '63mm' ? '1' : oldWidth === '126mm' ? '2' : oldWidth === '189mm' ? '3' : '4';
        }
    }
    var cardsNow = document.getElementsByClassName("card");
    var oneCol = [], twoCol = [], threeCol = [];
    for (var i = 0; i < cardsNow.length; i++) {
        var w = cardsNow[i].style.width;
        if (w === '63mm') oneCol.push(cardsNow[i]);
        else if (w === '126mm') twoCol.push(cardsNow[i]);
        else if (w === '189mm') threeCol.push(cardsNow[i]);
    }
    for (var i = 0; i < threeCol.length; i++) {
        document.body.appendChild(threeCol[i]);
        var level3 = threeCol[i].getElementsByClassName("number")[0].innerHTML;
        var j = 0;
        for (; j < oneCol.length; j++) {
            if (oneCol[j].getElementsByClassName("number")[0].innerHTML === level3) break;
        }
        var item = (j === oneCol.length) ? oneCol.shift() : oneCol.splice(j, 1)[0];
        document.body.appendChild(item);
    }
    for (var i = 0; i < twoCol.length; i++) {
        document.body.appendChild(twoCol[i]);
    }
    for (var i = 0; i < oneCol.length; i++) {
        document.body.appendChild(oneCol[i]);
    }
};
