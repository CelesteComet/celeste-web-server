/*Get me all bags that are tagged hot and sexy*/

select * from bag where bag.id in (select bag.id from bag
join bagtags on bag.id = bagtags.bagId
join tags on tags.id = bagtags.tagId 
where tags.name in ('hot', 'sexy')
group by bag.id
having COUNT(bag.id) = 2);
