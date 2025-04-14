// const k = 5; // Количество символов в алфавите: {A, C, G, T, N}

// // Класс, представляющий узел бора
// class BohrVertex {
//     constructor() {
//         this.next_vrtx = Array(k).fill(-1); // Массив переходов по символам
//         this.pat_num = -1;                  // Номер образца, если узел является концом образца
//         this.suff_link = -1;                // Суффиксная ссылка
//         this.auto_move = Array(k).fill(-1);  // Массив автоматических переходов
//         this.par = -1;                       // Номер родительского узла
//         this.suff_flink = -1;                // Ссылка на терминальную вершину
//         this.flag = false;                   // Флаг, указывающий, что узел является концом образца
//         this.parr = '';                      // Символ, по которому перешли в этот узел
//     }
// }

// let bohr = [];    // Бор (префиксное дерево)
// let pattern = []; // Список образцов

// // Преобразует символ в индекс (A -> 0, C -> 1, G -> 2, T -> 3, N -> 4)
// function charToIndex(c) {
//     switch (c) {
//         case 'A': return 0;
//         case 'C': return 1;
//         case 'G': return 2;
//         case 'T': return 3;
//         case 'N': return 4;
//         default:
//             console.log("Недопустимый символ");
//             process.exit(1);
//             return -1;
//     }
// }

// // Создает новый узел бора
// function makeBohrVertex(p, c) {
//     const v = new BohrVertex();
//     for (let i = 0; i < k; i++) {
//         v.next_vrtx[i] = -1;
//         v.auto_move[i] = -1;
//     }
//     v.flag = false;
//     v.suff_link = -1;
//     v.par = p;
//     v.parr = c;
//     v.suff_flink = -1;
//     return v;
// }

// // Инициализирует бор, добавляя корневой узел
// function bohrIni() {
//     bohr.push(makeBohrVertex(0, '$'));
// }

// // Добавляет строку в бор
// function addStringToBohr(s) {
//     let num = 0; // Начинаем с корневого узла
//     for (let i = 0; i < s.length; i++) {
//         const ch = charToIndex(s[i]); // Преобразуем символ в индекс
//         if (bohr[num].next_vrtx[ch] === -1) { // Если переход по символу отсутствует
//             bohr.push(makeBohrVertex(num, s[i])); // Создаем новый узел
//             bohr[num].next_vrtx[ch] = bohr.length - 1; // Добавляем переход
//         }
//         num = bohr[num].next_vrtx[ch]; // Переходим к следующему узлу
//     }
//     bohr[num].flag = true; // Помечаем узел как конец образца
//     pattern.push(s); // Добавляем образец в список
//     bohr[num].pat_num = pattern.length - 1; // Сохраняем номер образца
// }

// // Возвращает суффиксную ссылку для узла v
// function getSuffLink(v) {
//     if (bohr[v].suff_link === -1) { // Если суффиксная ссылка еще не вычислена
//         if (v === 0 || bohr[v].par === 0) { // Если узел корневой или его родитель корневой
//             bohr[v].suff_link = 0;
//         } else {
//             bohr[v].suff_link = getAutoMove(getSuffLink(bohr[v].par), bohr[v].parr);
//         }
//     }
//     return bohr[v].suff_link;
// }

// // Возвращает переход автомата для узла v и символа ch
// function getAutoMove(v, ch) {
//     const idx = charToIndex(ch); // Преобразуем символ в индекс
//     if (bohr[v].auto_move[idx] === -1) { // Если автоматический переход еще не вычислен
//         if (bohr[v].next_vrtx[idx] !== -1) { // Если есть прямой переход
//             bohr[v].auto_move[idx] = bohr[v].next_vrtx[idx];
//         } else {
//             if (v === 0) { // Если узел корневой
//                 bohr[v].auto_move[idx] = 0;
//             } else {
//                 bohr[v].auto_move[idx] = getAutoMove(getSuffLink(v), ch);
//             }
//         }
//     }
//     return bohr[v].auto_move[idx];
// }

// // Возвращает суффиксную ссылку на терминальную вершину
// function getSuffFlink(v) {
//     if (bohr[v].suff_flink === -1) { // Если суффиксная ссылка еще не вычислена
//         const u = getSuffLink(v);
//         if (u === 0) { // Если суффиксная ссылка ведет в корень
//             bohr[v].suff_flink = 0;
//         } else {
//             if (bohr[u].flag) { // Если узел является концом образца
//                 bohr[v].suff_flink = u;
//             } else {
//                 bohr[v].suff_flink = getSuffFlink(u); // двигаемся вверх по дереву
//             }
//         }
//     }
//     return bohr[v].suff_flink;
// }

// // Проверяет, является ли узел концом образца, и добавляет результат
// function check(v, i, res) {
//     for (let u = v; u !== 0; u = getSuffFlink(u)) {
//         if (bohr[u].flag) { // Если узел является концом образца
//             res.push([i - pattern[bohr[u].pat_num].length + 1, bohr[u].pat_num + 1]);
//         }
//     }
// }

// // Ищет все вхождения образцов в строке s
// function findAllPos(s) {
//     const res = [];
//     let u = 0; // Начинаем с корневого узла
//     for (let i = 0; i < s.length; i++) {
//         u = getAutoMove(u, s[i]); // Переходим по автоматической ссылке
//         check(u, i + 1, res);    // Проверяем, является ли узел концом образца
//     }
//     return res;
// }

// // Разделяет строку с учетом позиций разделителя
// function splitWithPositions(input, delimiter) {
//     const substrings = [];
//     const positions = [];
//     let buffer = [];
//     let ind = 0;

//     for (const r of input) {
//         if (r === delimiter) {
//             // Если текущий символ — разделитель, добавляем буфер в результат
//             if (buffer.length > 0) {
//                 substrings.push(buffer.join(''));
//                 positions.push(ind - buffer.length + 1);
//                 buffer = []; // Очищаем буфер
//             }
//         } else {
//             // Иначе добавляем символ в буфер
//             buffer.push(r);
//         }
//         ind += 1;
//     }
//     // Добавляем последний буфер, если он не пустой
//     if (buffer.length > 0) {
//         substrings.push(buffer.join(''));
//         positions.push(ind - buffer.length + 1);
//     }
//     return { substrings, positions };
// }

// // Основная функция
// function main() {
//     const args = process.argv.slice(2);
//     let numFlag = 1;
    
//     // Обработка флагов (